package prefecturejp_test

import (
	"testing"

	"github.com/theplant/prefecturejp"
	"github.com/theplant/testingutils/errorassert"
)

func TestGetPrefectures(t *testing.T) {
	prefectures := prefecturejp.GetPrefectures()

	if got, want := len(prefectures), 47; got != want {
		t.Fatalf("got %d of prefectures, but want %d", got, want)
	}
}

func TestGetPrefectureJapaneses(t *testing.T) {
	prefectures := prefecturejp.GetPrefectures()
	japaneses := prefecturejp.GetPrefectureJapaneses()

	for i, _ := range japaneses {
		if got, want := japaneses[i], prefectures[i].Japanese; got != want {
			t.Errorf("cases[%d] failed: got %q, but want %q", i, got, want)
		}
	}
}

func TestGetPrefectureNames(t *testing.T) {
	prefectures := prefecturejp.GetPrefectures()
	names := prefecturejp.GetPrefectureNames()

	for i, _ := range names {
		if got, want := names[i], prefectures[i].Name; got != want {
			t.Errorf("cases[%d] failed: got %q, but want %q", i, got, want)
		}
	}
}

func TestGetPrefectureShortNames(t *testing.T) {
	prefectures := prefecturejp.GetPrefectures()
	shortNames := prefecturejp.GetPrefectureShortNames()

	for i, _ := range shortNames {
		if got, want := shortNames[i], prefectures[i].ShortName; got != want {
			t.Errorf("cases[%d] failed: got %q, but want %q", i, got, want)
		}
	}
}

func TestGetPrefecturesJSON(t *testing.T) {
	prefectures, err := prefecturejp.GetPrefecturesJSON()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got, want := len(prefectures), 47; got != want {
		t.Fatalf("got %d of prefectures, but want %d", got, want)
	}
}

func TestGetJapaneseByName(t *testing.T) {
	var cases = []struct {
		name     string
		japanese string
	}{
		{"Hokkaido", "北海道"},
		{"Aomori", "青森県"},
		{"Iwate", "岩手県"},
		{"", ""},
	}

	for i, c := range cases {
		got := prefecturejp.GetJapaneseByName(c.name)
		if got != c.japanese {
			t.Errorf("cases[%d] failed: got %q, but want %q", i, got, c.japanese)
		}
	}
}

func TestGetNameByJapanese(t *testing.T) {
	var cases = []struct {
		name     string
		japanese string
	}{
		{"Hokkaido", "北海道"},
		{"Aomori", "青森県"},
		{"Iwate", "岩手県"},
		{"", ""},
	}

	for i, c := range cases {
		got := prefecturejp.GetNameByJapanese(c.japanese)
		if want := c.name; got != want {
			t.Errorf("cases[%d] failed: got %q, but want %q", i, got, want)
		}
	}
}

func TestGetShortNameByName(t *testing.T) {
	var cases = []struct {
		name      string
		shortName string
	}{
		{"Hokkaido", "北海道"},
		{"Aomori", "青森"},
		{"Iwate", "岩手"},
		{"", ""},
	}

	for i, c := range cases {
		got := prefecturejp.GetShortNameByName(c.name)
		if want := c.shortName; got != want {
			t.Errorf("cases[%d] failed: got %q, but want %q", i, got, want)
		}
	}
}

func TestGetNameByShortName(t *testing.T) {
	var cases = []struct {
		name      string
		shortName string
	}{
		{"Hokkaido", "北海道"},
		{"Aomori", "青森"},
		{"Iwate", "岩手"},
		{"", ""},
	}

	for i, c := range cases {
		got := prefecturejp.GetNameByShortName(c.shortName)
		if want := c.name; got != want {
			t.Errorf("cases[%d] failed: got %q, but want %q", i, got, want)
		}
	}
}

func TestGetCodeByJapanese(t *testing.T) {
	var cases = []struct {
		code     string
		japanese string
	}{
		{"JP-01", "北海道"},
		{"JP-02", "青森県"},
		{"JP-03", "岩手県"},
		{"", ""},
		{"", "xxx"},
	}

	for i, c := range cases {
		got := prefecturejp.GetCodeByJapanese(c.japanese)
		if want := c.code; got != want {
			t.Errorf("cases[%d] failed: got %q, but want %q", i, got, want)
		}
	}
}

func TestQueryPrefectureByFields(t *testing.T) {
	tests := []struct {
		name         string
		fields       []prefecturejp.PrefectureField
		value        string
		expectedPref *prefecturejp.Prefecture
	}{
		{
			name: "found",
			fields: []prefecturejp.PrefectureField{
				prefecturejp.PrefectureField_CODE,
				prefecturejp.PrefectureField_NAME,
				prefecturejp.PrefectureField_JAPANESE,
				prefecturejp.PrefectureField_SHORT_NAME,
			},
			value:        "Hokkaido",
			expectedPref: &prefecturejp.Prefecture{Code: "JP-01", Name: "Hokkaido", Japanese: "北海道", ShortName: "北海道", Region: "北海道"},
		},

		{
			name: "don't found",
			fields: []prefecturejp.PrefectureField{
				prefecturejp.PrefectureField_CODE,
			},
			value: "Aomori",
			expectedPref: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pref := prefecturejp.QueryPrefectureByFields(test.fields, test.value)
			errorassert.Equal(t, test.expectedPref, pref)
		})
	}
}
