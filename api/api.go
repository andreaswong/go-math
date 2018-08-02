package api

import "github.com/sirupsen/logrus"

type DTO interface{}

func EvenOdd(reqDTO DTO) (DTO, error) {
	logrus.Infof("req=%#v")
}
