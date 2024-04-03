package services

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"medcare/constants"
	"medcare/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mailgun/mailgun-go/v4"
)

func CreateToken(email, customerid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      email,
		"customerid": customerid, // Set the customerid claim
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(constants.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractID(jwtToken string, secretKey string) (string, error) {
	// Parse the JWT token
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err // Handle token parsing errors
	}

	// Check if the token is valid
	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Extract the customer ID from the claims
			customerID, ok := claims["customerid"].(string)
			if ok {
				return customerID, nil
			}
		}
	}

	return "", fmt.Errorf("invalid or expired JWT token")
}

func GenerateIDPA() string {
	length := 4
	randomBytes := make([]byte, length)
	// Use crypto/rand to generate random bytes
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err) // Handle error appropriately
	}
	// Convert random bytes to a hexadecimal string
	id := hex.EncodeToString(randomBytes)

	return "PAT" + id
}
func GenerateIDDOC() string {
	length := 4
	randomBytes := make([]byte, length)
	// Use crypto/rand to generate random bytes
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err) // Handle error appropriately
	}
	// Convert random bytes to a hexadecimal string
	id := hex.EncodeToString(randomBytes)

	return "DOC" + id
}
func GenerateIDBA() string {
	length := 4
	randomBytes := make([]byte, length)
	// Use crypto/rand to generate random bytes
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err) // Handle error appropriately
	}
	// Convert random bytes to a hexadecimal string
	id := hex.EncodeToString(randomBytes)

	return "BA" + id
}

func CurrentTime() string {
	currentTime := time.Now()
	currentTimeStr := currentTime.Format("2006-01-02 15:04:05")
	return currentTimeStr
}

// //////////////////////////// Send Email - mailGun //////////////////////////////////
func SendSimpleMessage(request *models.MailGunEmail, attachmentName string, attachmentData []byte) (string, error) {
	domain := "sandbox567bd12d5b42484e94273fbe408ababa.mailgun.org"
	mg := mailgun.NewMailgun(domain, constants.Mailgun_apikey)

	sender := "Excited User <mailgun@sandbox567bd12d5b42484e94273fbe408ababa.mailgun.org>"
	subject := request.Subject
	message := request.Message
	recipient := request.RecipientEmail // Ensure the recipient email address is correctly provided

	m := mg.NewMessage(sender, subject, message, recipient)

	// Add the attachment
	m.AddBufferAttachment(attachmentName, attachmentData)

	ctx := context.Background()

	// Create channels for tracking results and errors
	resultChan := make(chan string)
	errorChan := make(chan error)

	// Send emails concurrently
	go func() {
		// Send the message with the context
		resp, id, err := mg.Send(ctx, m)
		if err != nil {
			errorChan <- err
			return
		}
		fmt.Println("MailGun :", resp)
		resultChan <- id
	}()

	// Wait for results from goroutines
	select {
	case id := <-resultChan:
		return id, nil
	case err := <-errorChan:
		return "", err
	}
}

type MeetLinkResponse struct {
	HangoutLink string `json:"hangout_link"`
	EventID     string `json:"event_id"`
}

type Event struct {
	Summary        string    `json:"summary"`
	Start          time.Time `json:"start"`
	End            time.Time `json:"end"`
	ConferenceData `json:"conferenceData"`
	From           string `json:"from"`
	To             string `json:"to"`
}

type ConferenceData struct {
	CreateRequest `json:"createRequest"`
}

type CreateRequest struct {
	ConferenceSolutionKey `json:"conferenceSolutionKey"`
	RequestId             string `json:"requestId"`
}

type ConferenceSolutionKey struct {
	Type string `json:"type"`
}

// GetMeetLink function definition
// Update the GetMeetLink function to accept date, fromdatetime, and todatetime as inputs
func GetMeetLink(request *models.BookAppointment) (string, error, []byte) {
	url := "http://localhost:5001/create-event"

	// Event details
	startDateTime, err := time.Parse("2006-01-02 15:04:05", request.Date+" "+request.FromDateTime)
	if err != nil {
		return "", fmt.Errorf("error parsing start date and time: %w", err), nil
	}

	endDateTime, err := time.Parse("2006-01-02 15:04:05", request.Date+" "+request.ToDateTime)
	if err != nil {
		return "", fmt.Errorf("error parsing end date and time: %w", err), nil
	}

	from := request.DoctorEmail
	to := request.EmailID

	event := Event{
		Summary: "Medcare Appointment",
		Start:   startDateTime,
		End:     endDateTime,
		ConferenceData: ConferenceData{
			CreateRequest: CreateRequest{
				ConferenceSolutionKey: ConferenceSolutionKey{
					Type: "hangoutsMeet",
				},
				RequestId: "medcare.",
			},
		},
		From: from,
		To:   to,
	}

	// Convert event struct to JSON
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return "", fmt.Errorf("error marshalling event to JSON: %w", err), nil
	}

	// Make POST request to the microservice with event details
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(eventJSON))
	if err != nil {
		return "", fmt.Errorf("error making POST request: %w", err), nil
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response status: %d", resp.StatusCode), nil
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err), nil
	}

	// Parse the JSON response
	var meetLinkResponse MeetLinkResponse
	if err := json.Unmarshal(body, &meetLinkResponse); err != nil {
		return "", fmt.Errorf("error parsing JSON response: %w", err), nil
	}
	icsData, err := GenerateICalendarFile("Medcare Appointment", request.Date+" "+request.FromDateTime, request.Date+" "+request.ToDateTime, meetLinkResponse.HangoutLink)
	if err != nil {
		return "", fmt.Errorf("error generating iCalendar file: %w", err), nil
	}

	// Return the generated meet link
	return meetLinkResponse.HangoutLink, nil, icsData
}

func GenerateICalendarFile(summary, start, end, hangoutLink string) ([]byte, error) {
	startDateTime, err := time.Parse("2006-01-02 15:04:05", start)
	if err != nil {
		return nil, fmt.Errorf("error parsing start date and time: %w", err)
	}

	endDateTime, err := time.Parse("2006-01-02 15:04:05", end)
	if err != nil {
		return nil, fmt.Errorf("error parsing end date and time: %w", err)
	}

	eventID := uuid.New().String() // Generate a unique event ID
	event := fmt.Sprintf(`BEGIN:VCALENDAR
VERSION:2.0
BEGIN:VEVENT
UID:%s
SUMMARY:%s
DTSTART:%s
DTEND:%s
DESCRIPTION:Join the meeting using the link: %s
END:VEVENT
END:VCALENDAR`,
		eventID, summary, startDateTime.Format("20060102T150405"), endDateTime.Format("20060102T150405"), hangoutLink)

	return []byte(event), nil
}
