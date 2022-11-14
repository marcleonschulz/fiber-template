package config

import "testing"

func TestViper(t *testing.T) {
	var err error
	err = Init()
	if err != nil {
		t.Error(err)
	}
}
