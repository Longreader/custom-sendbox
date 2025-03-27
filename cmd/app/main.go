package app

import (
	"SandBox/internal/terminal"
	"github.com/sirupsen/logrus"
)

func Init() {
	logrus.Info("Terminal application init")
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: false,
		TimestampFormat:  "2006-01-02 15:04:05",
		FullTimestamp:    true,

		ForceColors:            true,
		DisableColors:          false,
		QuoteEmptyFields:       true,
		ForceQuote:             false,
		DisableLevelTruncation: true,
		PadLevelText:           true,
		// Customizing delimiters
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "[time]",
			logrus.FieldKeyLevel: "[level]",
			logrus.FieldKeyMsg:   "[content]",
		},
	})
	logrus.SetLevel(logrus.InfoLevel)
}

func Run() {
	var (
		term   *terminal.Terminal
		output string
		err    error
	)

	term = terminal.NewTerminal("cisco", "cisco", "172.16.255.35")

	_, err = term.GetExpectedPromptCiscoIOS()
	if err != nil {
		logrus.Fatal(err)
	}

	err = term.CiscoEnable("ciscoenable")
	if err != nil {
		logrus.Fatal(err)
	}

	output, err = term.WriteStdinExpected("sh run | i access-list", term.TermExpect)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("output: %s", output)
}
