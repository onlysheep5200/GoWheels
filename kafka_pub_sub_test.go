package GoWheels

import "testing"

func TestSub(t *testing.T) {
	subDoneChan := make(chan interface{})
	go sub(subDoneChan)

	pubDoneChan := make(chan interface{})
	go pub(pubDoneChan)

	<-subDoneChan
	<-pubDoneChan
}