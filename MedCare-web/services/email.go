package services

import (
	"bytes"
	"fmt"
	"medcare/models"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(subject string, content string, to []string, cc []string, bcc []string, attachmentName string, attachmentData []byte) error
}

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(name string, fromEmailAddress string, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *GmailSender) SendEmail(subject string, content string, to []string, cc []string, bcc []string, attachmentName string, attachmentData []byte) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	// Convert attachmentData byte slice to a bytes.Reader
	attachmentReader := bytes.NewReader(attachmentData)

	// Attach the ICS file
	_, err := e.Attach(attachmentReader, attachmentName, "text/calendar")
	if err != nil {
		return err
	}

	// Authenticate and send the email
	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)
}

var sender = NewGmailSender("MedCare.", "vishnusri196688@gmail.com", "xcjk utkc rnma exlq")

func SendEmailforAppointment(request *models.BookAppointment, attachmentName string, attachmentData []byte) error {
	subject := "Reg - Medcare Appointment"

	// Construct the HTML content for the email
	htmlTemplate := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
		    <meta charset="UTF-8">
		    <meta name="viewport" content="width=device-width, initial-scale=1.0">
		    <title>Medcare Appointment Confirmation</title>
		    <style>
		        /* Define your CSS styles here */
		        body {
		            font-family: Arial, sans-serif;
		            line-height: 1.6;
		            margin: 0;
		            padding: 0;
		            background-color: #f5f5f5;
		        }
		        .container {
		            max-width: 600px;
		            margin: 20px auto;
		            padding: 20px;
		            background-color: #fff;
		            border-radius: 12px;
		            box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
		            border: 1px solid #ddd;
		        }
		        h1 {
		            color: #333;
		            text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.2);
		        }
		        p {
		            color: #666;
		        }
		        ul {
		            list-style-type: none;
		            padding: 0;
		        }
		        li {
		            margin-bottom: 10px;
		        }
		        a {
		            color: #007bff;
		            text-decoration: none;
		        }
		        a:hover {
		            text-decoration: underline;
		        }
		        /* Add more styles as needed */
		    </style>
		</head>
		<body>
		    <div class="container">
		        <h1>Dear ` + request.Name + `, </h1>
		        <p>This is to confirm your appointment with ` + request.DoctorName + ` scheduled on ` + request.Date + ` at ` + request.FromDateTime + `.</p>
		        <h2>Appointment Details:</h2>
		        <ul>
		            <li>Doctor: ` + request.DoctorName + `</li>
		            <li>Date: ` + request.Date + `</li>
		            <li>Time: ` + request.FromDateTime + `</li>
		            <li>Virtual Meeting Link: <a href="` + request.MeetLink + `">Meet</a></li>
		        </ul>
		        <p>Please make sure to be available and ready for the appointment at least 5 minutes before the scheduled time.</p>
		        <p>If you have any questions or need to reschedule, please contact us at [Your Contact Information].</p>
		        <p>We look forward to seeing you.</p>
		        <p>Best regards,<br>MedCare</p>
		    </div>
		</body>
		</html>
	`

	// Email recipient
	to := []string{request.EmailID}
	fmt.Println(to)

	// Call SendEmail to send email with attachment
	err := sender.SendEmail(subject, htmlTemplate, to, nil, nil, attachmentName, attachmentData)
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	fmt.Println("Email sent successfully!")
	return nil
}
