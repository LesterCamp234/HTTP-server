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

type Product_json struct {
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

type Error struct {
	Error string
}

type Filtred_prducts struct {
	Title []string `json:"title"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Request-Headers", "*")

		content_type := r.Header.Get("Content-Type")

		if content_type == "application/json" {

			accept := r.Header.Get("Accept")

			w.Header().Set("Content-Type", accept)

			//0 = JSON | 1 = HTML
			var print_style int

			switch accept {
			case "text/html":
				print_style = 1
			case "application/json":
				print_style = 0
			}

			password := r.URL.Query().Get("password")

			file_id, err := strconv.Atoi(r.URL.Query().Get("file_id"))

			index, err_i := strconv.Atoi(r.URL.Query().Get("index"))

			if password == "pigna" {
				if err != nil && err_i != nil {
					if print_style == 1 {
						fmt.Fprintf(w, "<h1 style='text-align: center'> There is an error with the file_id and index </h1>")
					} else {
						err, _ := json.Marshal(Error{Error: "Invalid file_id and index"})
						fmt.Fprintf(w, "%s", err)

					}

				} else {
					Print_file(file_id, index, w, print_style)
				}
			} else {
				if print_style == 1 {
					fmt.Fprintf(w, "<h1 style='text-align: center'> Wrong password</h1>")
				} else {
					err, _ := json.Marshal(Error{Error: "Wrong password"})
					fmt.Fprintf(w, "%s", err)
				}
			}
		} else {
			w.Header().Set("Content-Type", "text/plain")
			fmt.Fprintf(w, "HTTP 406 - Not Acceptable")
		}

	})

	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Request-Headers", "*")

		content_type := r.Header.Get("Content-Type")

		if content_type == "application/json" {

			var min, max float64
			var err_min, err_max error

			accept := r.Header.Get("Accept")

			w.Header().Set("Content-Type", accept)

			//0 = JSON | 1 = HTML
			var print_style int

			switch accept {
			case "text/html":
				print_style = 1
			case "application/json":
				print_style = 0
			}

			if r.URL.Query().Has("min_rating") {
				min, err_min = strconv.ParseFloat(r.URL.Query().Get("min_rating"), 64)
			} else {
				min = 0
			}

			if r.URL.Query().Has("max_rating") {
				max, err_max = strconv.ParseFloat(r.URL.Query().Get("max_rating"), 64)
			} else {
				max = 5
			}

			if err_min != nil || err_max != nil {
				if print_style == 1 {
					fmt.Fprintf(w, "<h1 style='text-align: center'> There is an error with the ratings </h1>")
				} else {
					err, _ := json.Marshal(Error{Error: "Invalid ratings "})
					fmt.Fprintf(w, "%s", err)
				}

			} else {
				list_product(min, max, w, print_style)
			}
		} else {
			w.Header().Set("Content-Type", "text/plain")
			fmt.Fprintf(w, "HTTP 406 - Not Acceptable")
		}
	})

	http.ListenAndServe(":8080", nil)
}

func Print_file(file_id int, index int, w http.ResponseWriter, print_style int) {

	if file_id > 1 {
		index = index - ((file_id - 1) * 10)
	}

	index = index - 1

	if index > -1 && index < 10 {
		path := "test-files/test-" + strconv.Itoa(file_id) + ".json"
		file, err_file := os.ReadFile(path)

		if err_file != nil {
			if print_style == 1 {
				fmt.Fprintf(w, "<h1 style='text-align: center'>Unable to load the file</h1>")
			} else {
				err, _ := json.Marshal(Error{Error: "Unable to load the files"})
				fmt.Fprintf(w, "%s", err)
			}
		} else {
			var product Products
			err_file = json.Unmarshal(file, &product)
			if err_file != nil {
				if print_style == 1 {
					fmt.Fprintf(w, "<p>%s</p>", err_file)
				} else {
					err, _ := json.Marshal(Error{Error: err_file.Error()})
					fmt.Fprintf(w, "%s", err)
				}
			} else {
				if print_style == 1 {
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
				} else {
					product_json, _ := json.Marshal(Product_json{ID: product[index].ID, Title: product[index].Title, Description: product[index].Description, Price: product[index].Price,
						DiscountPercentage: product[index].DiscountPercentage, Rating: product[index].Rating, Stock: product[index].Stock, Brand: product[index].Brand,
						Category: product[index].Category, Thumbnail: product[index].Thumbnail, Images: product[index].Images})
					fmt.Fprintf(w, "%s", product_json)
				}
			}

		}
	} else {
		if print_style == 1 {
			fmt.Fprintf(w, "<h1 style='text-align: center'> There's an error with the file_id and index </h1>")
		} else {
			err, _ := json.Marshal(Error{Error: "There's an error with the file_id and index"})
			fmt.Fprintf(w, "%s", err)
		}
	}
}

func list_product(min float64, max float64, w http.ResponseWriter, print_style int) {

	var titles []string

	if min > max {
		if print_style == 1 {
			fmt.Fprintf(w, "<h1 style='text-align: center'> There is an error with the ratings </h1>")
		} else {
			err, _ := json.Marshal(Error{Error: "Invalid ratings"})
			fmt.Fprintf(w, "%s", err)
		}
	} else {

		for i := 1; i < 4; i++ {
			path := "test-files/test-" + strconv.Itoa(i) + ".json"
			file, err_file := os.ReadFile(path)

			if err_file != nil {
				if print_style == 1 {
					fmt.Fprintf(w, "<h1 style='text-align: center'>Unable to load the file</h1>")
				} else {
					err, _ := json.Marshal(Error{Error: "Unable to load the file"})
					fmt.Fprintf(w, "%s", err)
				}
			} else {
				var product Products
				err_file = json.Unmarshal(file, &product)
				if err_file != nil {
					if print_style == 1 {
						fmt.Fprintf(w, "<p>%s</p>", err_file)
					} else {
						err, _ := json.Marshal(Error{Error: err_file.Error()})
						fmt.Fprintf(w, "%s", err)
					}
				} else {
					if print_style == 1 {
						for j := 0; j < len(product); j++ {
							if product[j].Rating > min && product[j].Rating < max {
								fmt.Fprintf(w, "<p>%s</p>", product[j].Title)

							}
						}
					} else {
						for j := 0; j < len(product); j++ {
							if product[j].Rating > min && product[j].Rating < max {
								titles = append(titles, product[j].Title)
							}
						}
					}
				}
			}
		}
		if print_style != 1 {
			filtred_prducts, _ := json.Marshal(Filtred_prducts{Title: titles})
			fmt.Fprintf(w, "%s", filtred_prducts)
		}

	}

}
