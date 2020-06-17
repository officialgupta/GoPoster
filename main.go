package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"os"
	"io"
	"strconv"
	"io/ioutil"
) 

type Movies struct {
	Page         int `json:"page"`
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Results      []struct {
		Popularity       float64 `json:"popularity"`
		VoteCount        int     `json:"vote_count"`
		Video            bool    `json:"video"`
		PosterPath       string  `json:"poster_path"`
		ID               int     `json:"id"`
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		GenreIds         []int   `json:"genre_ids"`
		Title            string  `json:"title"`
		VoteAverage      int     `json:"vote_average"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
	} `json:"results"`
}

func main() {
	client := &http.Client{}
	var base = "https://image.tmdb.org/t/p/original/"

	var TMDBKey, err = ioutil.ReadFile("key")
    if err != nil {
        fmt.Print(err)
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("https://api.themoviedb.org/3/movie/popular?api_key=%s&language=en-GB&region=GB", TMDBKey), nil)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	
	var result Movies
	json.NewDecoder(resp.Body).Decode(&result)

	movies := result.Results

	for _,movie := range movies{
		fmt.Println(movie.Title)
		id := movie.ID
		URL := base + movie.PosterPath
		err := downloadFile(URL,strconv.Itoa(id)+".jpg")
		if err != nil {
			log.Fatal(err)
		}
	}

}

func downloadFile(URL, fileName string) error {
    response, err := http.Get(URL)
    if err != nil {
		log.Fatal(err)
    }
    defer response.Body.Close()

	path := "posters/"+fileName
	os.MkdirAll("posters", os.ModePerm)
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = io.Copy(file, response.Body)
    if err != nil {
        return err
    }
    return nil
}