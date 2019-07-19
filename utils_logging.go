package main

import logrus "github.com/sirupsen/logrus"

var logger *logrus.Logger

type fields map[string]interface{}

func loggerCreate() error {
	logger = logrus.New()

	return nil
}
