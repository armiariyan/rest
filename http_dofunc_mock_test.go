package rest

import (
	"testing"
)

func TestDoFuncHttpErr(t *testing.T) {
	_, err := DoFuncHttpErr(nil)
	if err == nil {
		t.Error("DoFuncHttpErr must return not nil error")
		t.Fail()
	}
}

func TestDoFuncErrParseBodyResp(t *testing.T) {
	_, err := DoFuncErrParseBodyResp(nil)
	if err != nil {
		t.Error("DoFuncErrParseBodyResp should return nil")
		t.Fail()
	}
}
