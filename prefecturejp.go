package prefecturejp

// nameToJapaneseMap maps the `Name` to `Japanese`
var nameToJapaneseMap = map[string]string{}

// nameToShortNameMap maps the `Name` to `ShortName`
var nameToShortNameMap = map[string]string{}

// japaneseToNameMap maps the `Japanese` to `Name`.
var japaneseToNameMap = map[string]string{}

// shortNameToNameMap maps the `ShortName` to `Name`.
var shortNameToNameMap = map[string]string{}

// Prefectures contains a list of prefecture.
type Prefectures []Prefecture

// Prefecture which defines codes for the names of the principal subdivisions
// (e.g.provinces or states) of all countries coded in ISO 3166-1.
//
// See: https://en.wikipedia.org/wiki/ISO_3166-2:JP
type Prefecture struct {
	// The standard code, like: JP-01, JP-02, etc
	Code string `json:"code"`
	// The english name, like: Hokkaido, Aomori etc
	Name string `json:"name"`
	// The japanese name, like: 北海道, 岩手県 etc
	Japanese string `json:"japanese"`
	// The short name of japanese, like: 北海道, 青森 etc.
	ShortName string `json:"short_name"`
	Region    string `json:"region"`
}

func init() {
	for _, prefecture := range prefectureData {
		nameToJapaneseMap[prefecture.Name] = prefecture.Japanese
		nameToShortNameMap[prefecture.Name] = prefecture.ShortName
		japaneseToNameMap[prefecture.Japanese] = prefecture.Name
		shortNameToNameMap[prefecture.ShortName] = prefecture.Name
	}
}

// GetPrefectures return the Japan prefecture list.
func GetPrefectures() Prefectures {
	return prefectureData
}

// GetPrefectureJapaneses return all the `Japanese`s of the Prefectures.
func GetPrefectureJapaneses() []string {
	prefectureJapaneses := []string{}
	for _, prefecture := range prefectureData {
		prefectureJapaneses = append(prefectureJapaneses, prefecture.Japanese)
	}
	return prefectureJapaneses
}

// GetPrefectureNames return the all the `Name`s of the Prefectures.
func GetPrefectureNames() []string {
	prefectureNames := []string{}
	for _, prefecture := range prefectureData {
		prefectureNames = append(prefectureNames, prefecture.Name)
	}
	return prefectureNames
}

// GetPrefectureShortNames return all the `ShortName`s of the Prefectures.
func GetPrefectureShortNames() []string {
	prefectureShortNames := []string{}
	for _, prefecture := range prefectureData {
		prefectureShortNames = append(prefectureShortNames, prefecture.ShortName)
	}
	return prefectureShortNames
}

// GetPrefecturesJSON is a wrapper helper to jsonhandlerfunc.
func GetPrefecturesJSON() (Prefectures, error) {
	return GetPrefectures(), nil
}

// GetJapaneseByName returns the `Japanese` by the given `Name`.
func GetJapaneseByName(name string) (result string) {
	return nameToJapaneseMap[name]
}

// GetShortNameByName returns the `ShortName` by the given `Name`.
func GetShortNameByName(name string) (result string) {
	return nameToShortNameMap[name]
}

// GetNameByJapanese returns the `Name` by the given `Japanese`.
func GetNameByJapanese(japanese string) (result string) {
	return japaneseToNameMap[japanese]
}

// GetNameByShortName returns the `Name` by the given `ShortName`.
func GetNameByShortName(shortName string) (result string) {
	return shortNameToNameMap[shortName]
}
