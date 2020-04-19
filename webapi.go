package datasetgen

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const OpenDataSiteDomain string = "opendata.pref.aomori.lg.jp"
const OpenDataWebApiUrl string = "https://" + OpenDataSiteDomain + "/api/package_show?id=5e4612ce-1636-41d9-82a3-c5130a79ffe0"

type WebApi struct {
	Success bool `json:"success"`
	Result  struct {
		Type string `json:"type"`
		Site struct {
			ID      int       `json:"id"`
			Name    string    `json:"name"`
			URL     string    `json:"url"`
			Created time.Time `json:"created"`
			Updated time.Time `json:"updated"`
		} `json:"site"`
		User struct {
			ID      int       `json:"id"`
			Name    string    `json:"name"`
			UID     string    `json:"uid"`
			Email   string    `json:"email"`
			Created time.Time `json:"created"`
			Updated time.Time `json:"updated"`
		} `json:"user"`
		Author     string    `json:"author"`
		ID         int       `json:"id"`
		UUID       string    `json:"uuid"`
		Name       string    `json:"name"`
		Filename   string    `json:"filename"`
		URL        string    `json:"url"`
		Text       string    `json:"text"`
		State      string    `json:"state"`
		Released   time.Time `json:"released"`
		Created    time.Time `json:"created"`
		Updated    time.Time `json:"updated"`
		Categories []struct {
			ID       int       `json:"id"`
			Name     string    `json:"name"`
			Filename string    `json:"filename"`
			State    string    `json:"state"`
			Created  time.Time `json:"created"`
			Updated  time.Time `json:"updated"`
		} `json:"categories"`
		EstatCategories []interface{} `json:"estat_categories"`
		Areas           []struct {
			ID       int       `json:"id"`
			Name     string    `json:"name"`
			Filename string    `json:"filename"`
			State    string    `json:"state"`
			Created  time.Time `json:"created"`
			Updated  time.Time `json:"updated"`
		} `json:"areas"`
		Tags       []interface{} `json:"tags"`
		Point      int           `json:"point"`
		Downloaded int           `json:"downloaded"`
		Groups     []struct {
			ID           int       `json:"id"`
			Name         string    `json:"name"`
			TrailingName string    `json:"trailing_name"`
			Created      time.Time `json:"created"`
			Updated      time.Time `json:"updated"`
		} `json:"groups"`
		Resources []struct {
			ID         int         `json:"id"`
			UUID       string      `json:"uuid"`
			RevisionID string      `json:"revision_id"`
			Name       string      `json:"name"`
			Filename   string      `json:"filename"`
			Text       interface{} `json:"text"`
			License    struct {
				ID      int         `json:"id"`
				Name    string      `json:"name"`
				UID     interface{} `json:"uid"`
				State   string      `json:"state"`
				Created time.Time   `json:"created"`
				Updated time.Time   `json:"updated"`
			} `json:"license"`
			RdfIri      interface{} `json:"rdf_iri"`
			RdfError    interface{} `json:"rdf_error"`
			Created     time.Time   `json:"created"`
			Updated     time.Time   `json:"updated"`
			DownloadURL string      `json:"download_url"`
			URL         string      `json:"url"`
			Format      string      `json:"format"`
		} `json:"resources"`
		NumResources int `json:"num_resources"`
	} `json:"result"`
	Help string `json:"help"`
}

func NewWebApi() *WebApi {
	return new(WebApi)
}

func (w *WebApi) Get() error {
	resp, err := http.Get(OpenDataWebApiUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteArray, &w)
	return err
}
