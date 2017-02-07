




* [Get Japanese By Name](#get-japanese-by-name)
* [Get Name By Japanese](#get-name-by-japanese)
* [Get Name By Short Name](#get-name-by-short-name)
* [Get Prefecture Japaneses](#get-prefecture-japaneses)
* [Get Prefecture Names](#get-prefecture-names)
* [Get Prefecture Short Names](#get-prefecture-short-names)
* [Get Short Name By Name](#get-short-name-by-name)
* [Type Prefecture](#type-prefecture)
* [Type Prefectures](#type-prefectures)
  * [Get Prefectures](#prefectures-get-prefectures)
  * [Get Prefectures JS ON](#prefectures-get-prefectures-js-on)




## Get Japanese By Name
``` go
func GetJapaneseByName(name string) (result string)
```
GetJapaneseByName returns the `Japanese` by the given `Name`.



## Get Name By Japanese
``` go
func GetNameByJapanese(japanese string) (result string)
```
GetNameByJapanese returns the `Name` by the given `Japanese`.



## Get Name By Short Name
``` go
func GetNameByShortName(shortName string) (result string)
```
GetNameByShortName returns the `Name` by the given `ShortName`.



## Get Prefecture Japaneses
``` go
func GetPrefectureJapaneses() []string
```
GetPrefectureJapaneses return all the `Japanese`s of the Prefectures.



## Get Prefecture Names
``` go
func GetPrefectureNames() []string
```
GetPrefectureNames return the all the `Name`s of the Prefectures.



## Get Prefecture Short Names
``` go
func GetPrefectureShortNames() []string
```
GetPrefectureShortNames return all the `ShortName`s of the Prefectures.



## Get Short Name By Name
``` go
func GetShortNameByName(name string) (result string)
```
GetShortNameByName returns the `ShortName` by the given `Name`.





## Type: Prefecture
``` go
type Prefecture struct {
    // The standard code, like: JP-01, JP-02, etc
    Code string `json:"code"`
    // The english name, like: Hokkaido, Aomori etc
    Name string `json:"name"`
    // The japanese name, like: 北海道, 岩手県 etc
    Japanese string `json:"japanese"`
    // The short name of japanese, like: 北海道, 青森 etc.
    ShortName string `json:"short_name"`
}
```
Prefecture which defines codes for the names of the principal subdivisions
(e.g.provinces or states) of all countries coded in ISO 3166-1.

See: <a href="https://en.wikipedia.org/wiki/ISO_3166-2:JP">https://en.wikipedia.org/wiki/ISO_3166-2:JP</a>










## Type: Prefectures
``` go
type Prefectures []Prefecture
```
Prefectures contains a list of prefecture.







### Prefectures: Get Prefectures
``` go
func GetPrefectures() Prefectures
```
GetPrefectures return the Japan prefecture list.


### Prefectures: Get Prefectures JS ON
``` go
func GetPrefecturesJSON() (Prefectures, error)
```
GetPrefecturesJSON is a wrapper helper to jsonhandlerfunc.






