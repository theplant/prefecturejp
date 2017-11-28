package prefecturejp

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

var codeIndexMap = map[string]int{}
var nameIndexMap = map[string]int{}
var japaneseIndexMap = map[string]int{}
var shortNameIndexMap = map[string]int{}
var regionIndexMap = map[string]int{}

func init() {
	for i, prefecture := range prefectureData {
		codeIndexMap[prefecture.Code] = i
		nameIndexMap[prefecture.Name] = i
		japaneseIndexMap[prefecture.Japanese] = i
		shortNameIndexMap[prefecture.ShortName] = i
		regionIndexMap[prefecture.Region] = i
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
	prefecture := QueryPrefecture(PrefectureField_NAME, name)
	if prefecture == nil {
		return ""
	}
	return prefecture.Japanese
}

// GetShortNameByName returns the `ShortName` by the given `Name`.
func GetShortNameByName(name string) (result string) {
	prefecture := QueryPrefecture(PrefectureField_NAME, name)
	if prefecture == nil {
		return ""
	}
	return prefecture.ShortName
}

// GetNameByJapanese returns the `Name` by the given `Japanese`.
func GetNameByJapanese(japanese string) (result string) {
	prefecture := QueryPrefecture(PrefectureField_JAPANESE, japanese)
	if prefecture == nil {
		return ""
	}
	return prefecture.Name
}

// GetNameByShortName returns the `Name` by the given `ShortName`.
func GetNameByShortName(shortName string) (result string) {
	prefecture := QueryPrefecture(PrefectureField_SHORT_NAME, shortName)
	if prefecture == nil {
		return ""
	}
	return prefecture.Name
}

// GetCodeByJapanese returns the `Code` by given `Japanese`
func GetCodeByJapanese(japanese string) (result string) {
	prefecture := QueryPrefecture(PrefectureField_JAPANESE, japanese)
	if prefecture == nil {
		return ""
	}
	return prefecture.Code
}

func QueryPrefecture(field PrefectureField, value string) *Prefecture {
	var prefectureIndex int
	var isFound bool
	switch field {
	case PrefectureField_CODE:
		prefectureIndex, isFound = codeIndexMap[value]
	case PrefectureField_NAME:
		prefectureIndex, isFound = nameIndexMap[value]
	case PrefectureField_JAPANESE:
		prefectureIndex, isFound = japaneseIndexMap[value]
	case PrefectureField_SHORT_NAME:
		prefectureIndex, isFound = shortNameIndexMap[value]
	case PrefectureField_REGION:
		prefectureIndex, isFound = regionIndexMap[value]
	}

	if !isFound {
		return nil
	}

	newPrefecture := prefectureData[prefectureIndex]
	return &newPrefecture
}
