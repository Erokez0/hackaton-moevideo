package skydns

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Erokez0/hackaton-moevideo/src/database"
	"maps"
)

type SkydnsModule struct {
	categories      []map[string]map[string]string
	categories_list map[string]string
}

var Skydns SkydnsModule = SkydnsModule{
	categories_list: make(map[string]string),
}

func (s *SkydnsModule) getSkydnsCategories() {
	response, err := http.Get("https://z.api.skydns.ru/catgroups")
	if err != nil {
		panic("Error getting Skydns categories")
	}
	defer response.Body.Close()
	result, err := io.ReadAll(response.Body)
	if err != nil {
		panic("Error reading Skydns categories")
	}
	json.Unmarshal(result, &s.categories)
}

func (s *SkydnsModule) getSkydnsCategoriesMap() {
	for _, category := range s.categories {
		for _, cat := range category {
			maps.Copy(s.categories_list, cat)
		}
	}
}

func (s *SkydnsModule) SkydnsCategoryIdtoName(id string) string {
	if name, ok := s.categories_list[id]; ok {
		return name
	}
	return ""
}

func (s *SkydnsModule) CategorizerSkydns(requestUrl string) string {
	url := "https://z.api.skydns.ru/domain/" + requestUrl
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer response.Body.Close()
	var category struct {
		Category []int
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	json.Unmarshal(bytes, &category)
	id := strconv.Itoa(category.Category[0])

	result := s.SkydnsCategoryIdtoName(id)
	return result
}

func Categorize(requestUrl string, confident bool) []int {
	categoryName := Skydns.CategorizerSkydns(requestUrl)
	// log.Println(categoryName)
	if categoryName == "" {
		return []int{}
	}
	return database.FindIdsLikeName(categoryName, confident)
}

func Init() {
	Skydns.getSkydnsCategories()
	Skydns.getSkydnsCategoriesMap()
	log.Println("\x1b[32mSkydns categorizer initialized\x1b[0m")
}
