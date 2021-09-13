package main

import (
	"testing"
)

const (
	empty_str = ""
	simple_error_str = "qwerty.smerty"
	simple_work_url = "https://open.spotify.com/"
	long_str = "aspndasngopiwddfjsdfapsdoeriuetpasdfpnsdgpnsdpfisdpfihsdpfhasfhasofhasofhas[fhashfas[fh[spfhasdhfoashfsfhjlsfhsljkfhlkjshflkjshfkljshfkjshfkjashfkjashfjasfhkahfsfowutbfadpubfaosdufhsdufhwuehrbdf;jasbdg;asjdbgls;dfhwoeurhioduf[iuwritpwiqiosjglkajglaskjgapsgihpiwhtpwihtpiwt"
)

func TestAdd(t *testing.T) {
	if _, err := Shorting(); err != nil { t.Errorf("Shorting error: %v\n", err) }
	if _, err := Shorting(); err != nil { t.Errorf("Shorting error: %v\n", err) }
	if _, err := Shorting(); err != nil { t.Errorf("Shorting error: %v\n", err) }
	if _, err := Shorting(); err != nil { t.Errorf("Shorting error: %v\n", err) }
	if err := Validate(empty_str); err == nil { t.Errorf("Validate error: %v\n", err) }
	if err := Validate(simple_error_str); err == nil { t.Errorf("Validate error: %v\n", err) }
	if err := Validate(simple_work_url); err != nil { t.Errorf("Validate error: %v\n", err) }
	if err := Validate(long_str); err == nil { t.Errorf("Validate error: %v\n", err) }
}

