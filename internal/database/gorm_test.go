package database

import "testing"

func TestGorm(t *testing.T) {
	var err error
	err = Init()
	if err != nil {
		t.Error(err)
	}
	err = HealthCheck()
	if err != nil {
		t.Error(err)
	}
}
