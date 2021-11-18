package model

import (
	"testing"
)

func TestNew(t *testing.T) {
	newCar := New("M2", "BMW", "2020")
	if newCar.Key == "" || len(newCar.Key) != 36{
		t.Error("func New Car não gerou Key com 36 posições para novo veículo")
	}
}