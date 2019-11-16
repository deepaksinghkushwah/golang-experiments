package urlshort

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

//MapHandler return
func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

//YAMLHandler return nil nil
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// 1. parse the yaml
	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}

	// 2. convert yaml array to map
	pathToUrls := buildMap(pathUrls)
	// 3. return a MapHandler using map
	return MapHandler(pathToUrls, fallback), nil
}

func buildMap(pathUrls []pathURL) map[string]string {
	pathToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathToUrls[pu.Path] = pu.URL
	}
	return pathToUrls
}

func parseYaml(data []byte) ([]pathURL, error) {
	var pathUrls []pathURL
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
