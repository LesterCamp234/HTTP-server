package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Products []struct {
	ID                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Price              int      `json:"price"`
	DiscountPercentage float64  `json:"discountPercentage"`
	Rating             float64  `json:"rating"`
	Stock              int      `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
	Thumbnail          string   `json:"thumbnail"`
	Images             []string `json:"images"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		password := r.URL.Query().Get("password")

		file_id, err := strconv.Atoi(r.URL.Query().Get("file_id"))

		index, err_i := strconv.Atoi(r.URL.Query().Get("index"))

		if password == "pigna" {
			if err != nil && err_i != nil {
				fmt.Fprintf(w, "<h1 style='text-align: center'> There is an error with the file_id and index </h1>")
			} else {
				Print_file(file_id, index, w)
			}
		} else {
			fmt.Fprintf(w, "<h1 style='text-align: center'> Wrong password</h1>")
		}

	})

	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {

		var min, max float64
		var err_min, err_max error

		if r.URL.Query().Has("min_rating") {
			min, err_min = strconv.ParseFloat(r.URL.Query().Get("min_rating"), 64)
		} else {
			min = 0
		}

		if r.URL.Query().Has("max_rating") {
			max, err_max = strconv.ParseFloat(r.URL.Query().Get("min_rating"), 64)
		} else {
			max = 5
		}

		if err_min != nil || err_max != nil {
			fmt.Fprintf(w, "<h1 style='text-align: center'> There is an error with the ratings </h1>")
		} else {
			list_product(min, max, w)
		}

	})

	http.ListenAndServe(":8080", nil)
}

func Print_file(file_id int, index int, w http.ResponseWriter) {

	if file_id > 1 {
		index = index - ((file_id - 1) * 10)
	}

	index = index - 1

	if index > -1 && index < 10 {
		path := "test-files/test-" + strconv.Itoa(file_id) + ".json"
		file, err_file := os.ReadFile(path)

		if err_file != nil {
			fmt.Fprintf(w, "<h1 style='text-align: center>Unable to load the file</h1>")
		} else {

			var product Products
			err_file = json.Unmarshal(file, &product)
			if err_file != nil {
				fmt.Fprintf(w, "<p>%s</p>", err_file)
			} else {
				fmt.Fprintf(w, "<h1 style='text-align: center'>%s</h1>", product[index].Title)
				fmt.Fprintf(w, "<img src='%s'>", product[index].Thumbnail)
				fmt.Fprintf(w, "<p>%s</p>", product[index].Description)
				fmt.Fprintf(w, "<p>%d</p>", product[index].Price)
				fmt.Fprintf(w, "<p>%f</p>", product[index].DiscountPercentage)
				fmt.Fprintf(w, "<p>%f</p>", product[index].Rating)
				fmt.Fprintf(w, "<p>%d</p>", product[index].Stock)
				fmt.Fprintf(w, "<p>%s</p>", product[index].Brand)
				fmt.Fprintf(w, "<div style='flex: auto'>")
				for i := 0; i < len(product[index].Images); i++ {
					fmt.Fprintf(w, "<img src='%s' width='200' style='padding-left: 20px'>", product[index].Images[i])
				}
				fmt.Fprintf(w, "</div>")
			}

		}
	} else {
		fmt.Fprintf(w, "<h1 style='text-align: center'> There's an error with the file_id and index </h1>")
	}
}

func list_product(min float64, max float64, w http.ResponseWriter) {

	if min > max {
		fmt.Fprintf(w, "<h1 style='text-align: center'> There is an error with the ratings </h1>")
	} else {
		for i := 1; i < 4; i++ {
			path := "test-files/test-" + strconv.Itoa(i) + ".json"
			file, err_file := os.ReadFile(path)

			if err_file != nil {
				fmt.Fprintf(w, "<h1 style='text-align: center>Unable to load the file</h1>")
			} else {
				var product Products
				err_file = json.Unmarshal(file, &product)
				if err_file != nil {
					fmt.Fprintf(w, "<p>%s</p>", err_file)
				} else {
					for j := 0; j < len(product); j++ {
						if product[j].Rating > min && product[j].Rating < max {
							fmt.Fprintf(w, "<p>%s</p>", product[j].Title)
						}
					}
				}
			}
		}

	}

}
