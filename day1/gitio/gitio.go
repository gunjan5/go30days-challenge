// Package gitio is a client for http://git.io URL shortener.
package gitio

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
)

const (
	gitioPostAPI = "http://git.io/create"
	gitioPutAPI  = "http://git.io"
	gitioGetAPI  = "http://git.io"
)

const gitioAPI = "http://git.io/create"

// Shorten returns a short version of an URL, or an error otherwise.
// Please note that it's not guaranteed the code will be accepted by git.io,
// the random one may be used instead.
func Shorten(longURL) (shortURL *url.URL, err error) {
	if len(longURL) == 0 {
		return nil, errors.New("no URL provided")
	}

	form := make(url.Values)
	form.Add("url", longURL)

	var api string
	api = gitioPostAPI

	resp, err := http.PostForm(api, form)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusCreated: // for PUT
		return resp.Location()
	case http.StatusOK: // for POST
		randomCode, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		} else if len(randomCode) == 0 {
			return nil, errors.New("unknown error")
		}
		u, _ := url.Parse(gitioGetAPI)
		u.Path = filepath.Join(u.Path, string(randomCode))
		return u, nil
	case http.StatusInternalServerError, 422:
		return nil, errors.New("only GitGub/Gist links are accepted")
	default:
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}
}

// CheckTaken checks if the provided custom code has already been taken on git.io.
func CheckTaken(code string) (bool, error) {
	if len(code) == 0 {
		return false, errors.New("no code provided")
	}
	u, _ := url.Parse(gitioGetAPI)
	u.Path = filepath.Join(u.Path, code)
	resp, err := http.Get(u.String())
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusNotFound:
		return false, nil
	case http.StatusFound:
		return true, nil
	default:
		// probably it's taken
		return true, nil
	}
}
