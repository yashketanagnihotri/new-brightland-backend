package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

type FormData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Query string `json:"query"`
}

func main() {
	http.HandleFunc("/submit-query", handleQuery)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var data FormData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = sendEmail(data)
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Query received and email sent!"})
}

func sendEmail(data FormData) error {
	// Replace with your SMTP server details
	from := "yashagni1992@gmail.com"
	password := "zvcs khen mngc uqdm"
	to := "yashagni1992@gmail.com"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "New Query from " + data.Name
	body := fmt.Sprintf("Name: %s\nEmail: %s\nPhone: %s\nQuery: %s",
		data.Name, data.Email, data.Phone, data.Query)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" + body

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	return err
}

/*

{
  "name": "Aarav Sharma",
  "email": "aarav@example.com",
  "phone": "9876543210",
  "query": "What are the admission fees and timings for preschool?"
}


*/
