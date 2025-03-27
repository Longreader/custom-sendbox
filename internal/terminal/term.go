package terminal

import (
	"SandBox/internal/crypto/ssh"
	"bufio"
	"errors"
	"fmt"
	//"golang.org/x/crypto/ssh"
	"io"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var keyExchanges []string = []string{
	"diffie-hellman-group-exchange-sha256",
	"diffie-hellman-group14-sha1",
	"diffie-hellman-group-exchange-sha1",
	"diffie-hellman-group1-sha1",
}

var (
	CiscoASAInput = regexp.MustCompile(".*Invalid input detected at '\\^' marker")
)

//TODO: Связка с PTY
//const (
//	terminalWidth  = 511
//	terminalHeight = 256
//	terminalSpeed  = 14400
//	terminalType   = "vt100"
//)

const (
	ExpectedAny  = ".*"
	ExpectedSpec = ".*[>#].*"
)

const (
	ErrTerminalTimeout = "timeout expired"
)

type Terminal struct {
	TermExpect string
	client     *ssh.Client
	session    *ssh.Session
	stdIn      io.WriteCloser
	stdOut     io.Reader
	stdErr     io.Reader
	errChan    chan error
}

// NewTerminal returns new terminal instance for by_ssh communication
func NewTerminal(login, password, host string) *Terminal {
	var (
		sshConfig    *ssh.Config
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session

		err error
	)

	sshConfig = &ssh.Config{
		Ciphers:        []string{"aes128-cbc"},
		RekeyThreshold: 256 * 1024,
	}
	sshConfig.SetDefaults()
	sshConfig.KeyExchanges = append(sshConfig.KeyExchanges, keyExchanges...)

	clientConfig = &ssh.ClientConfig{
		Config:          *sshConfig,
		User:            login,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	client, err = ssh.Dial("tcp", fmt.Sprintf("%s:22", host), clientConfig)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("failed to dial")
		return nil
	}

	session, err = client.NewSession()
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("failed to create session")
		_ = client.Close()
		return nil
	}

	logrus.Debug("terminal allocated")

	return &Terminal{
		client:  client,
		session: session,
		errChan: make(chan error, 1),
	}
}

func (t *Terminal) Run() error {
	var (
		//modes by_ssh.TerminalModes
		err error
	)

	//TODO: Необходимо ли настраивать PTY?
	//modes = by_ssh.TerminalModes{
	//	by_ssh.ECHO:          0,             // Отключаем эхо
	//	by_ssh.ICRNL:         0,             // Корректная обработка Enter
	//	by_ssh.ONLCR:         0,             // Перевод \n в \r\n
	//	by_ssh.TTY_OP_ISPEED: terminalSpeed, // input speed = 14.4 KK
	//	by_ssh.TTY_OP_OSPEED: terminalSpeed, // output speed = 14.4 KK
	//}
	//
	//err = t.session.RequestPty(terminalType, terminalHeight, terminalWidth, modes)
	//if err != nil {
	//	logrus.WithField("error", err).Error("failed to request terminal")
	//	if err = t.Close(); err != nil {
	//		logrus.WithField("error", err).Error("failed to close terminal")
	//		return err
	//	}
	//	return err
	//}

	t.stdIn, err = t.session.StdinPipe()
	if err != nil {
		logrus.WithField("error", err).Error("failed to set stdin for terminal")
		if err = t.Close(); err != nil {
			logrus.WithField("error", err).Error("failed to close terminal")
			return err
		}
		return err
	}

	t.stdOut, err = t.session.StdoutPipe()
	if err != nil {
		logrus.WithField("error", err).Error("failed to set stdout for terminal")
		if err = t.Close(); err != nil {
			logrus.WithField("error", err).Error("failed to close terminal")
			return err
		}
		return err
	}

	t.stdErr, err = t.session.StderrPipe()
	if err != nil {
		logrus.WithField("error", err).Error("failed to set stderr for terminal")
		if err = t.Close(); err != nil {
			logrus.WithField("error", err).Error("failed to close terminal")
			return err
		}
		return err
	}

	err = t.session.Shell()
	if err != nil {
		logrus.WithField("error", err).Error("failed to start shell")
		if err = t.Close(); err != nil {
			logrus.WithField("error", err).Error("failed to close terminal")
			return err
		}
		return err
	}

	logrus.Debug("terminal started")

	return nil
}

func (t *Terminal) Close() error {
	var err error

	err = t.session.Close()
	if err != nil && !errors.Is(err, os.ErrClosed) && !errors.Is(err, io.EOF) {
		logrus.WithFields(logrus.Fields{"error": err}).Error("failed to close session")
		return err
	}
	return nil
}

func (t *Terminal) StartReading() {
	go func() {
		scanner := bufio.NewScanner(t.stdOut)
		for scanner.Scan() {
			logrus.Debug("[OUTPUT] ", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			logrus.WithField("error", err).Error("failed to read stdout")
		}
	}()
}

func (t *Terminal) getStdout(expected string, timer time.Duration) string {
	var (
		scanner *bufio.Scanner
		regExp  *regexp.Regexp
		local   string
		output  string
	)
	scanner = bufio.NewScanner(t.stdOut)
	regExp = regexp.MustCompile(fmt.Sprintf(`(?m)^%s`, expected))
	ticker := time.NewTicker(timer)
	defer ticker.Stop()
	for scanner.Scan() {
		local = scanner.Text()
		local = strings.Trim(local, "\b")
		local = strings.Trim(local, "\r")
		logrus.Debug("[EXPECTED PRINT] ", local)
		output += local + "\n"
		if regExp.MatchString(local) {
			logrus.WithFields(logrus.Fields{"" +
				"expected": expected,
				"actual": local,
			}).Warning("[OUTPUT] detached")
			scanner = nil
			return output
		}
		select {
		case <-ticker.C:
			return ""
		default:
			continue
		}
	}
	if scanner.Err() != nil {
		logrus.WithField("error", scanner.Err()).Error("failed to read stdout")
	}
	return ""
}

func (t *Terminal) ReadStdoutExpected(timeout time.Duration, expected string) string {
	outputChan := make(chan string, 1)

	go func() {
		outputChan <- t.getStdout(expected, timeout)
	}()

	select {
	case output := <-outputChan:
		return output
	case <-time.After(timeout):
		logrus.Warning("Timeout while waiting for output")
		return ErrTerminalTimeout
	}
}

func (t *Terminal) WriteStdinExpected(cmd string, expected string) (string, error) {
	var (
		err error
	)

	_, err = fmt.Fprintf(t.stdIn, cmd+"\n\r")
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("failed to send input")
		if err = t.Close(); err != nil {
			logrus.WithFields(logrus.Fields{"error": err}).Error("failed to close terminal")
			return "", err
		}
		return "", err
	}

	logrus.WithFields(logrus.Fields{"cmd": cmd}).Debug("send input")

	out := t.ReadStdoutExpected(25*time.Second, expected)
	if out == ErrTerminalTimeout {
		return "", errors.New(ErrTerminalTimeout)
	}

	var output strings.Builder
	regCmd := regexp.MustCompile(cmd)
	regPrompt := regexp.MustCompile(expected)

	for _, line := range strings.Split(out, "\n") {
		if regCmd.MatchString(line) || regPrompt.MatchString(line) {
			continue
		}
		output.WriteString(line + "\n")
	}

	logrus.WithFields(logrus.Fields{
		"cmd":      cmd,
		"expected": expected,
		"output":   output.String(),
	}).Debug("get output")

	return output.String(), nil
}

func (t *Terminal) Wait() error {
	err := <-t.errChan
	return err
}
