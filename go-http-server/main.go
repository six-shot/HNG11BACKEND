package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

type IPInfo struct {
	IP       string `json:"client_ip"`
	Location string `json:"location"`
	Greeting string `json:"greeting"`
	
}

func getIPInfo() (*IPInfo, error) {
	

	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

   city := data["city"].(string)

   apiKey := os.Getenv("OPENWEATHER_API_KEY")
	weatherResp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey))
	if err != nil {
		return nil, err
	}
	defer weatherResp.Body.Close()

	weatherBody, err := io.ReadAll(weatherResp.Body)
	if err != nil {
		return nil, err
	}
	var weatherData map[string]interface{}
	if err := json.Unmarshal(weatherBody, &weatherData); err != nil {
		return nil, err
	}

	
	temperature := weatherData["main"].(map[string]interface{})["temp"].(float64)

	name := os.Getenv("NAME")
	if name == "" {
		name = "six-shot"
	}

		greeting := fmt.Sprintf("Hello, %s!, The temperature is %.1f degrees Celsius in %s", name, temperature, city)

	info := &IPInfo{
		IP:       data["ip"].(string),
		Location: data["city"].(string), // Using "city" as Location
		Greeting: greeting,
		
	}

	return info, nil
}




func handler(w http.ResponseWriter, r *http.Request) {
	ipInfo, err := getIPInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ipInfo)
}

func main() {
	
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
