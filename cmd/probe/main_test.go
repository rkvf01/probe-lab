package main

import (
	"testing"
)

func TestParseTargets(t *testing.T) {

	testcases := map[string]struct {
		input   string
		wantErr bool
		wantLen int
	}{
		"duplicate names": {input: `[{"name":"api","url":"http://x"},{"name":"api","url":"http://xyz"} ]`, wantErr: true},
		"valid":           {input: `[{"name":"api","url":"http://x"},{"name":"auth","url":"http://xyz"} ]`, wantLen: 2},
		"empty_array":     {input: `[]`, wantLen: 0},
		"malformed":       {input: `[{bad`, wantErr: true},
		"missing url":     {input: `[{"name":"api"}]`, wantErr: true},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			targets, err := parseTargets([]byte(tc.input))

			if (err != nil) != tc.wantErr {
				t.Errorf("FAIL: Got error %v , expected error %v ", err, tc.wantErr)
				return
			}
			if tc.wantErr {
				return // the error was expected; nothing more to check
			}

			if len(targets) != tc.wantLen {
				t.Errorf("FAIL: Got %d targets, expected %d targets", len(targets), tc.wantLen)
				return
			}

		})
	}

}
