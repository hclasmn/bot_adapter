package coolq

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestDeCode(t *testing.T) {
	str:="1111111[CQ:face,id=115,text=111][CQ:face,id=217]] [CQ:text,text=123] ["
	a:=DeCode(str)
	log.Infof("map:%+v",a)
}
