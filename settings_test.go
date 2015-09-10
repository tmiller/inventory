package main

import "testing"

func Test_loadSettings(t *testing.T) {

	expected := settings{
		ServerPort:      "8080",
		ForemanUsername: "foo",
		ForemanPassword: "bar",
	}

	actual := loadSettings("./test/good_settings.json")

	if expected.ServerPort != actual.ServerPort {
		t.Errorf("ServerPort Expected: '%s', Actual:'%s'",
			expected.ServerPort,
			actual.ServerPort,
		)
	}

	if expected.ForemanUsername != actual.ForemanUsername {
		t.Errorf("ForemanUsername Expected: '%s', Actual:'%s'",
			expected.ForemanUsername,
			actual.ForemanUsername,
		)
	}

	if expected.ForemanPassword != actual.ForemanPassword {
		t.Errorf("ForemanPassword Expected: '%s', Actual:'%s'",
			expected.ForemanPassword,
			actual.ForemanPassword,
		)
	}
}
