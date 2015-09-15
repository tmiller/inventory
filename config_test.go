package main

import "testing"

func Test_loadConfig(t *testing.T) {

	expected := config{
		ServerPort:      "8080",
		ForemanUsername: "foo",
		ForemanPassword: "bar",
		ForemanURL:      "http://localhost",
	}

	actual, err := loadConfig("./testdata/good_config.json")

	if err != nil {
		t.Error(err)
	}

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

	if expected.ForemanURL != actual.ForemanURL {
		t.Errorf("ForemanPassword Expected: '%s', Actual:'%s'",
			expected.ForemanURL,
			actual.ForemanURL,
		)
	}
}
