package terminal

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func (t *Terminal) GetExpectedPromptCiscoASA() (string, error) {
	var (
		err error
	)
	_, err = fmt.Fprintf(t.stdIn, "\n\r")
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("failed to send input")
		if err = t.Close(); err != nil {
			logrus.WithFields(logrus.Fields{"error": err}).Error("failed to close terminal")
			return "", err
		}

		return "", err
	}

	output := t.ReadStdoutExpected(time.Second*25, ".*[>#]")
	if output == ErrTerminalTimeout {
		return "", errors.New(ErrTerminalTimeout)
	}

	output = strings.Replace(output, "\n\r", "\n", -1)
	lines := strings.Split(output, "\n")
	lastLine := strings.TrimSpace(lines[len(lines)-2])
	lastLine = strings.Replace(lastLine, ">", "", -1)
	lastLine = strings.Replace(lastLine, "#", "", -1)
	expectedPrompt := fmt.Sprintf("%s.*[>#]", lastLine)

	t.TermExpect = expectedPrompt

	return expectedPrompt, err
}

func (t *Terminal) GetExpectedPromptCiscoIOS() (string, error) {
	var (
		err error
	)
	_, err = fmt.Fprintf(t.stdIn, "\n\r")
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("failed to send input")
		if err = t.Close(); err != nil {
			logrus.WithFields(logrus.Fields{"error": err}).Error("failed to close terminal")
			return "", err
		}

		return "", err
	}

	output := t.ReadStdoutExpected(time.Second*25, ".*[>#]")
	if output == ErrTerminalTimeout {
		return "", errors.New(ErrTerminalTimeout)
	}

	output = strings.Replace(output, "\n\r", "\n", -1)
	lines := strings.Split(output, "\n")
	lastLine := strings.TrimSpace(lines[len(lines)-2])
	lastLine = strings.Replace(lastLine, ">", "", -1)
	lastLine = strings.Replace(lastLine, "#", "", -1)
	expectedPrompt := fmt.Sprintf("%s.*[>#]", lastLine)

	t.TermExpect = expectedPrompt

	return expectedPrompt, err
}

func (t *Terminal) CiscoEnable(enablePass string) error {
	var (
		err error
	)
	_, err = t.stdIn.Write([]byte("enable" + "\n\r"))
	_, err = t.stdIn.Write([]byte(enablePass + "\n\r"))
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("failed to send input")
		if err = t.Close(); err != nil {
			logrus.WithFields(logrus.Fields{"error": err}).Error("failed to close terminal")
			return err
		}

		return err
	}
	logrus.WithFields(logrus.Fields{"cmd": "enable"}).Debug("send input")

	output := t.ReadStdoutExpected(time.Second*25, ".*word")
	if output == ErrTerminalTimeout {
		return errors.New(ErrTerminalTimeout)
	}

	return nil
}

func (t *Terminal) NoPagerCiscoASA(expectedPrompt string) error {
	var (
		err error
	)
	_, err = fmt.Fprintf(t.stdIn, "terminal pager 0"+"\n\r")
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("failed to send input")
		if err = t.Close(); err != nil {
			logrus.WithFields(logrus.Fields{"error": err}).Error("failed to close terminal")
			return err
		}
		return err
	}
	logrus.WithFields(logrus.Fields{"cmd": "terminal pager 0"}).Debug("send input")

	output := t.ReadStdoutExpected(time.Second*25, expectedPrompt)
	if output == ErrTerminalTimeout {
		return errors.New(ErrTerminalTimeout)
	}

	return nil
}

func (t *Terminal) NoPagerCiscoIOS(expectedPrompt string) error {
	var (
		err error
	)
	_, err = t.stdIn.Write([]byte("terminal length 0" + "\n\r"))
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("failed to send input")
		if err = t.Close(); err != nil {
			logrus.WithFields(logrus.Fields{"error": err}).Error("failed to close terminal")
			return err
		}
		return err
	}
	logrus.WithFields(logrus.Fields{"cmd": "terminal length 0"}).Debug("send input")

	output := t.ReadStdoutExpected(time.Second*25, expectedPrompt)
	if output == ErrTerminalTimeout {
		return errors.New(ErrTerminalTimeout)
	}

	return nil
}
