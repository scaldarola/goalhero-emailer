package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type BetaRegisterRequest struct {
	Email    string `json:"email"`
	Language string `json:"language"`
}

type BetaRegisterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(BetaRegisterResponse{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	var req BetaRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BetaRegisterResponse{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	if req.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BetaRegisterResponse{
			Success: false,
			Message: "Email is required",
		})
		return
	}

	if req.Language == "" {
		req.Language = "en"
	}

	if req.Language != "en" && req.Language != "es" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BetaRegisterResponse{
			Success: false,
			Message: "Language must be 'en' or 'es'",
		})
		return
	}

	if err := sendWelcomeEmail(req.Email, req.Language); err != nil {
		log.Printf("Error sending email: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(BetaRegisterResponse{
			Success: false,
			Message: "Failed to send welcome email",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(BetaRegisterResponse{
		Success: true,
		Message: "Welcome email sent successfully!",
	})
}

func sendWelcomeEmail(email, language string) error {
	from := mail.NewEmail("GoalHero Team", "info@goalhero.eu")
	to := mail.NewEmail("", email)

	var subject string
	var htmlContent string

	if language == "es" {
		subject = "🎉⚽ ¡Bienvenido a GoalHero!"
		htmlContent = getWelcomeEmailHTMLSpanish()
	} else {
		subject = "🎉⚽ Welcome to GoalHero!"
		htmlContent = getWelcomeEmailHTML()
	}

	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		fmt.Printf("Error sending email: %v\n", err)
		return fmt.Errorf("error sending email: %v", err)
	}

	if response.StatusCode >= 400 {
		fmt.Printf("Error sending email: %v\n", err)
		return fmt.Errorf("sendgrid error: status code %d", response.StatusCode)
	}

	return nil
}

func getWelcomeEmailHTML() string {
	return `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Welcome to GoalHero!</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            line-height: 1.6;
            color: #333;
            background-color: #f8fafc;
        }
        
        .container {
            max-width: 600px;
            margin: 0 auto;
            background-color: #ffffff;
            border-radius: 16px;
            overflow: hidden;
            box-shadow: 0 10px 40px rgba(0, 0, 0, 0.1);
        }
        
        .logo-banner {
            width: 100%;
            height: 200px;
            background: #000000;
            display: flex;
            align-items: center;
            justify-content: center;
            position: relative;
            padding: 20px;
        }
        
        .logo-banner img {
            max-width: 300px;
            max-height: 160px;
            width: auto;
            height: auto;
            display: block;
            margin: 0 auto;
        }
        
        .header {
            background: #ffffff;
            padding: 30px 30px 20px;
            text-align: center;
            border-bottom: 1px solid #e5e7eb;
        }
        
        .header h1 {
            color: #1a1a1a;
            font-size: 32px;
            font-weight: 700;
            margin-bottom: 10px;
        }
        
        .header p {
            color: #4a4a4a;
            font-size: 18px;
            font-weight: 500;
        }
        
        .content {
            padding: 40px 30px;
        }
        
        .welcome-message {
            text-align: center;
            margin-bottom: 40px;
        }
        
        .welcome-message h2 {
            font-size: 28px;
            color: #1a1a1a;
            margin-bottom: 20px;
            font-weight: 600;
        }
        
        .welcome-message p {
            font-size: 18px;
            color: #4a4a4a;
            line-height: 1.7;
            max-width: 500px;
            margin: 0 auto;
        }
        
        .features {
            background: linear-gradient(135deg, #f8f8f8 0%, #f0f0f0 100%);
            border-radius: 16px;
            padding: 30px;
            margin: 40px 0;
            border: 1px solid #e0e0e0;
        }
        
        .features h3 {
            font-size: 22px;
            color: #1a1a1a;
            margin-bottom: 25px;
            text-align: center;
            font-weight: 600;
        }
        
        .feature-list {
            list-style: none;
        }
        
        .feature-list li {
            padding: 12px 0;
            color: #2a2a2a;
            display: flex;
            align-items: center;
            font-size: 16px;
            font-weight: 500;
        }
        
        .feature-list li:before {
            content: "🥅";
            margin-right: 15px;
            font-size: 18px;
        }
        
        .cta-section {
            text-align: center;
            margin: 40px 0;
        }
        
        .cta-button {
            display: inline-block;
            background: linear-gradient(135deg, #00C851 0%, #007E33 100%);
            color: #ffffff;
            text-decoration: none;
            padding: 20px 50px;
            border-radius: 50px;
            font-weight: 700;
            font-size: 19px;
            letter-spacing: 0.5px;
            text-transform: uppercase;
            transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
            box-shadow: 0 8px 30px rgba(0, 200, 81, 0.4);
            border: 3px solid transparent;
            position: relative;
            overflow: hidden;
        }
        
        .cta-button:before {
            content: '';
            position: absolute;
            top: 0;
            left: -100%;
            width: 100%;
            height: 100%;
            background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
            transition: left 0.6s;
        }
        
        .cta-button:hover {
            transform: translateY(-5px) scale(1.05);
            box-shadow: 0 15px 40px rgba(0, 200, 81, 0.6);
            border-color: rgba(255, 255, 255, 0.3);
        }
        
        .cta-button:hover:before {
            left: 100%;
        }
        
        .cta-button:active {
            transform: translateY(-2px) scale(1.02);
            transition: all 0.1s ease;
        }
        
        .social-section {
            background: linear-gradient(135deg, #f8f8f8 0%, #f0f0f0 100%);
            border-radius: 16px;
            padding: 35px;
            text-align: center;
            margin: 40px 0;
            border: 1px solid #e0e0e0;
        }
        
        .social-section h3 {
            font-size: 22px;
            color: #1a1a1a;
            margin-bottom: 30px;
            font-weight: 600;
        }
        
        .social-links {
            display: flex;
            justify-content: center;
            align-items: center;
            gap: 20px;
            flex-wrap: wrap;
        }
        
        .social-link {
            display: inline-flex;
            align-items: center;
            justify-content: center;
            text-decoration: none;
            font-weight: 600;
            padding: 14px 28px;
            border-radius: 50px;
            transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
            font-size: 16px;
            min-width: 160px;
            position: relative;
            overflow: hidden;
        }
        
        .social-link.website {
            background: linear-gradient(135deg, #1a73e8 0%, #1557b0 100%);
            color: #ffffff;
            box-shadow: 0 6px 20px rgba(26, 115, 232, 0.3);
            border: 2px solid transparent;
        }
        
        .social-link.instagram {
            background: linear-gradient(45deg, #f09433 0%,#e6683c 25%,#dc2743 50%,#cc2366 75%,#bc1888 100%);
            color: #ffffff;
            box-shadow: 0 6px 20px rgba(225, 48, 108, 0.3);
            border: 2px solid transparent;
        }
        
        .social-link:before {
            content: '';
            position: absolute;
            top: 0;
            left: -100%;
            width: 100%;
            height: 100%;
            background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
            transition: left 0.6s;
        }
        
        .social-link:hover {
            transform: translateY(-3px) scale(1.05);
            box-shadow: 0 12px 35px rgba(0, 0, 0, 0.2);
        }
        
        .social-link.website:hover {
            box-shadow: 0 12px 35px rgba(26, 115, 232, 0.4);
        }
        
        .social-link.instagram:hover {
            box-shadow: 0 12px 35px rgba(225, 48, 108, 0.4);
        }
        
        .social-link:hover:before {
            left: 100%;
        }
        
        .social-link:active {
            transform: translateY(-1px) scale(1.02);
            transition: all 0.1s ease;
        }
        
        .footer {
            background: linear-gradient(135deg, #1a1a1a 0%, #000000 100%);
            color: #cccccc;
            padding: 40px 30px;
            text-align: center;
        }
        
        .footer p {
            margin-bottom: 15px;
            font-size: 16px;
        }
        
        .footer a {
            color: #4CAF50;
            text-decoration: none;
            transition: color 0.3s ease;
        }
        
        .footer a:hover {
            color: #66BB6A;
        }
        
        @media (max-width: 600px) {
            .container {
                margin: 10px;
                border-radius: 12px;
            }
            
            .logo-banner {
                height: 150px;
            }
            
            .logo-banner img {
                max-width: 200px;
                max-height: 80px;
            }
            
            .header, .content {
                padding: 25px 20px;
            }
            
            .header h1 {
                font-size: 28px;
            }
            
            .welcome-message h2 {
                font-size: 24px;
            }
            
            .welcome-message p {
                font-size: 16px;
            }
            
            .features, .social-section {
                padding: 25px 20px;
            }
            
            .social-links {
                flex-direction: column;
                align-items: center;
                gap: 10px;
            }
            
            .social-link {
                width: 200px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="logo-banner">
            <img src="https://www.goalhero.eu/assets/icon.png" alt="GoalHero Logo" />
        </div>
        
        <div class="header">
            <h1>Welcome to GoalHero!</h1>
            <p>Never cancel another match - find goalkeepers instantly</p>
        </div>
        
        <div class="content">
            <div class="welcome-message">
                <h2>⚽ Welcome to the Beta!</h2>
                <p>Thank you for joining GoalHero! You're among the first to experience our revolutionary goalkeeper marketplace that ensures your team never forfeits another match due to missing keepers.</p>
                <p style="margin-top: 20px; font-weight: 600; color: #00C851;">📱 We'll contact you as soon as the beta is ready for download!</p>
            </div>
            
            <div class="features">
                <h3>What's Coming Your Way</h3>
                <ul class="feature-list">
                    <li>Post games and receive competitive bids from goalkeepers</li>
                    <li>Browse verified goalkeeper profiles with ratings & reviews</li>
                    <li>Secure payment system with guaranteed show-up protection</li>
                </ul>
            </div>        
            
            <div class="social-section">
                <h3>Stay Connected</h3>
                <div class="social-links">
                    <a href="https://www.goalhero.eu" class="social-link website">
                        <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor" style="margin-right: 8px;">
                            <path d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm6.93 6h-2.95c-.32-1.25-.78-2.45-1.38-3.56 1.84.63 3.37 1.91 4.33 3.56zM12 4.04c.83 1.2 1.48 2.53 1.91 3.96h-3.82c.43-1.43 1.08-2.76 1.91-3.96zM4.26 14C4.1 13.36 4 12.69 4 12s.1-1.36.26-2h3.38c-.08.66-.14 1.32-.14 2 0 .68.06 1.34.14 2H4.26zm.82 2h2.95c.32 1.25.78 2.45 1.38 3.56-1.84-.63-3.37-1.9-4.33-3.56zm2.95-8H5.08c.96-1.66 2.49-2.93 4.33-3.56C8.81 5.55 8.35 6.75 8.03 8zM12 19.96c-.83-1.2-1.48-2.53-1.91-3.96h3.82c-.43 1.43-1.08 2.76-1.91 3.96zM14.34 14H9.66c-.09-.66-.16-1.32-.16-2 0-.68.07-1.35.16-2h4.68c.09.65.16 1.32.16 2 0 .68-.07 1.34-.16 2zm.25 5.56c.6-1.11 1.06-2.31 1.38-3.56h2.95c-.96 1.65-2.49 2.93-4.33 3.56zM16.36 14c.08-.66.14-1.32.14-2 0-.68-.06-1.34-.14-2h3.38c.16.64.26 1.31.26 2s-.1 1.36-.26 2h-3.38z"/>
                        </svg>
                        Website
                    </a>
                    <a href="https://instagram.com/goalhero.app" class="social-link instagram">
                        <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor" style="margin-right: 8px;">
                            <path d="M7.8 2h8.4C19.4 2 22 4.6 22 7.8v8.4a5.8 5.8 0 0 1-5.8 5.8H7.8C4.6 22 2 19.4 2 16.2V7.8A5.8 5.8 0 0 1 7.8 2m-.2 2A3.6 3.6 0 0 0 4 7.6v8.8C4 18.39 5.61 20 7.6 20h8.8a3.6 3.6 0 0 0 3.6-3.6V7.6C20 5.61 18.39 4 16.4 4H7.6m9.65 1.5a1.25 1.25 0 0 1 1.25 1.25A1.25 1.25 0 0 1 17.25 8 1.25 1.25 0 0 1 16 6.75a1.25 1.25 0 0 1 1.65-1.25M12 7a5 5 0 0 1 5 5 5 5 0 0 1-5 5 5 5 0 0 1-5-5 5 5 0 0 1 5-5m0 2a3 3 0 0 0-3 3 3 3 0 0 0 3 3 3 3 0 0 0 3-3 3 3 0 0 0-3-3z"/>
                        </svg>
                        Instagram
                    </a>
                </div>
            </div>
            
            <div style="text-align: center; margin-top: 40px; padding-top: 30px; border-top: 1px solid #e5e7eb;">
                <p style="color: #6b7280; font-size: 16px;">
                    Have questions? We're here to help! Reply to this email or contact us at 
                    <a href="mailto:info@goalhero.eu" style="color: #4CAF50;">info@goalhero.eu</a>
                </p>
            </div>
        </div>
        
        <div class="footer">
            <p><strong>GoalHero Team</strong></p>
            <p>Making dreams achievable, one goal at a time.</p>
            <p style="margin-top: 20px; font-size: 14px; opacity: 0.8;">
                © 2025 GoalHero. All rights reserved.<br>
                <a href="#">Unsubscribe</a> | <a href="#">Privacy Policy</a> | <a href="#">Terms of Service</a>
            </p>
        </div>
    </div>
</body>
</html>
    `
}

func getWelcomeEmailHTMLSpanish() string {
	return `
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>¡Bienvenido a GoalHero!</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            line-height: 1.6;
            color: #333;
            background-color: #f8fafc;
        }
        
        .container {
            max-width: 600px;
            margin: 0 auto;
            background-color: #ffffff;
            border-radius: 16px;
            overflow: hidden;
            box-shadow: 0 10px 40px rgba(0, 0, 0, 0.1);
        }
        
        .logo-banner {
            width: 100%;
            height: 200px;
            background: #000000;
            display: flex;
            align-items: center;
            justify-content: center;
            position: relative;
            padding: 20px;
        }
        
        .logo-banner img {
            max-width: 300px;
            max-height: 160px;
            width: auto;
            height: auto;
            display: block;
            margin: 0 auto;
        }
        
        .header {
            background: #ffffff;
            padding: 30px 30px 20px;
            text-align: center;
            border-bottom: 1px solid #e5e7eb;
        }
        
        .header h1 {
            color: #1a1a1a;
            font-size: 32px;
            font-weight: 700;
            margin-bottom: 10px;
        }
        
        .header p {
            color: #4a4a4a;
            font-size: 18px;
            font-weight: 500;
        }
        
        .content {
            padding: 40px 30px;
        }
        
        .welcome-message {
            text-align: center;
            margin-bottom: 40px;
        }
        
        .welcome-message h2 {
            font-size: 28px;
            color: #1a1a1a;
            margin-bottom: 20px;
            font-weight: 600;
        }
        
        .welcome-message p {
            font-size: 18px;
            color: #4a4a4a;
            line-height: 1.7;
            max-width: 500px;
            margin: 0 auto;
        }
        
        .features {
            background: linear-gradient(135deg, #f8f8f8 0%, #f0f0f0 100%);
            border-radius: 16px;
            padding: 30px;
            margin: 40px 0;
            border: 1px solid #e0e0e0;
        }
        
        .features h3 {
            font-size: 22px;
            color: #1a1a1a;
            margin-bottom: 25px;
            text-align: center;
            font-weight: 600;
        }
        
        .feature-list {
            list-style: none;
        }
        
        .feature-list li {
            padding: 12px 0;
            color: #2a2a2a;
            display: flex;
            align-items: center;
            font-size: 16px;
            font-weight: 500;
        }
        
        .feature-list li:before {
            content: "🥅";
            margin-right: 15px;
            font-size: 18px;
        }
        
        .cta-section {
            text-align: center;
            margin: 40px 0;
        }
        
        .cta-button {
            display: inline-block;
            background: linear-gradient(135deg, #00C851 0%, #007E33 100%);
            color: #ffffff;
            text-decoration: none;
            padding: 20px 50px;
            border-radius: 50px;
            font-weight: 700;
            font-size: 19px;
            letter-spacing: 0.5px;
            text-transform: uppercase;
            transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
            box-shadow: 0 8px 30px rgba(0, 200, 81, 0.4);
            border: 3px solid transparent;
            position: relative;
            overflow: hidden;
        }
        
        .cta-button:before {
            content: '';
            position: absolute;
            top: 0;
            left: -100%;
            width: 100%;
            height: 100%;
            background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
            transition: left 0.6s;
        }
        
        .cta-button:hover {
            transform: translateY(-5px) scale(1.05);
            box-shadow: 0 15px 40px rgba(0, 200, 81, 0.6);
            border-color: rgba(255, 255, 255, 0.3);
        }
        
        .cta-button:hover:before {
            left: 100%;
        }
        
        .cta-button:active {
            transform: translateY(-2px) scale(1.02);
            transition: all 0.1s ease;
        }
        
        .social-section {
            background: linear-gradient(135deg, #f8f8f8 0%, #f0f0f0 100%);
            border-radius: 16px;
            padding: 35px;
            text-align: center;
            margin: 40px 0;
            border: 1px solid #e0e0e0;
        }
        
        .social-section h3 {
            font-size: 22px;
            color: #1a1a1a;
            margin-bottom: 30px;
            font-weight: 600;
        }
        
        .social-links {
            display: flex;
            justify-content: center;
            align-items: center;
            gap: 20px;
            flex-wrap: wrap;
        }
        
        .social-link {
            display: inline-flex;
            align-items: center;
            justify-content: center;
            text-decoration: none;
            font-weight: 600;
            padding: 14px 28px;
            border-radius: 50px;
            transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
            font-size: 16px;
            min-width: 160px;
            position: relative;
            overflow: hidden;
        }
        
        .social-link.website {
            background: linear-gradient(135deg, #1a73e8 0%, #1557b0 100%);
            color: #ffffff;
            box-shadow: 0 6px 20px rgba(26, 115, 232, 0.3);
            border: 2px solid transparent;
        }
        
        .social-link.instagram {
            background: linear-gradient(45deg, #f09433 0%,#e6683c 25%,#dc2743 50%,#cc2366 75%,#bc1888 100%);
            color: #ffffff;
            box-shadow: 0 6px 20px rgba(225, 48, 108, 0.3);
            border: 2px solid transparent;
        }
        
        .social-link:before {
            content: '';
            position: absolute;
            top: 0;
            left: -100%;
            width: 100%;
            height: 100%;
            background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
            transition: left 0.6s;
        }
        
        .social-link:hover {
            transform: translateY(-3px) scale(1.05);
            box-shadow: 0 12px 35px rgba(0, 0, 0, 0.2);
        }
        
        .social-link.website:hover {
            box-shadow: 0 12px 35px rgba(26, 115, 232, 0.4);
        }
        
        .social-link.instagram:hover {
            box-shadow: 0 12px 35px rgba(225, 48, 108, 0.4);
        }
        
        .social-link:hover:before {
            left: 100%;
        }
        
        .social-link:active {
            transform: translateY(-1px) scale(1.02);
            transition: all 0.1s ease;
        }
        
        .footer {
            background: linear-gradient(135deg, #1a1a1a 0%, #000000 100%);
            color: #cccccc;
            padding: 40px 30px;
            text-align: center;
        }
        
        .footer p {
            margin-bottom: 15px;
            font-size: 16px;
        }
        
        .footer a {
            color: #4CAF50;
            text-decoration: none;
            transition: color 0.3s ease;
        }
        
        .footer a:hover {
            color: #66BB6A;
        }
        
        @media (max-width: 600px) {
            .container {
                margin: 10px;
                border-radius: 12px;
            }
            
            .logo-banner {
                height: 150px;
            }
            
            .logo-banner img {
                max-width: 200px;
                max-height: 80px;
            }
            
            .header, .content {
                padding: 25px 20px;
            }
            
            .header h1 {
                font-size: 28px;
            }
            
            .welcome-message h2 {
                font-size: 24px;
            }
            
            .welcome-message p {
                font-size: 16px;
            }
            
            .features, .social-section {
                padding: 25px 20px;
            }
            
            .social-links {
                flex-direction: column;
                align-items: center;
                gap: 10px;
            }
            
            .social-link {
                width: 200px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="logo-banner">
            <img src="https://www.goalhero.eu/assets/icon.png" alt="GoalHero Logo" />
        </div>
        
        <div class="header">
            <h1>¡Bienvenido a GoalHero!</h1>
            <p>Nunca canceles otro partido - encuentra porteros al instante</p>
        </div>
        
        <div class="content">
            <div class="welcome-message">
                <h2>⚽ ¡Bienvenido a la Beta!</h2>
                <p>¡Gracias por unirte a GoalHero! Estás entre los primeros en experimentar nuestro revolucionario marketplace de porteros que asegura que tu equipo nunca más tenga que abandonar un partido por falta de porteros.</p>
                <p style="margin-top: 20px; font-weight: 600; color: #00C851;">📱 ¡Te contactaremos tan pronto como la beta esté lista para descargar!</p>
            </div>
            
            <div class="features">
                <h3>Lo Que Te Espera</h3>
                <ul class="feature-list">
                    <li>Publica partidos y recibe ofertas competitivas de porteros</li>
                    <li>Explora perfiles verificados de porteros con calificaciones y reseñas</li>
                    <li>Sistema de pago seguro con protección de asistencia garantizada</li>
                </ul>
            </div>
        
            <div class="social-section">
                <h3>Mantente Conectado</h3>
                <div class="social-links">
                    <a href="https://www.goalhero.eu" class="social-link website">
                        <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor" style="margin-right: 8px;">
                            <path d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm6.93 6h-2.95c-.32-1.25-.78-2.45-1.38-3.56 1.84.63 3.37 1.91 4.33 3.56zM12 4.04c.83 1.2 1.48 2.53 1.91 3.96h-3.82c.43-1.43 1.08-2.76 1.91-3.96zM4.26 14C4.1 13.36 4 12.69 4 12s.1-1.36.26-2h3.38c-.08.66-.14 1.32-.14 2 0 .68.06 1.34.14 2H4.26zm.82 2h2.95c.32 1.25.78 2.45 1.38 3.56-1.84-.63-3.37-1.9-4.33-3.56zm2.95-8H5.08c.96-1.66 2.49-2.93 4.33-3.56C8.81 5.55 8.35 6.75 8.03 8zM12 19.96c-.83-1.2-1.48-2.53-1.91-3.96h3.82c-.43 1.43-1.08 2.76-1.91 3.96zM14.34 14H9.66c-.09-.66-.16-1.32-.16-2 0-.68.07-1.35.16-2h4.68c.09.65.16 1.32.16 2 0 .68-.07 1.34-.16 2zm.25 5.56c.6-1.11 1.06-2.31 1.38-3.56h2.95c-.96 1.65-2.49 2.93-4.33 3.56zM16.36 14c.08-.66.14-1.32.14-2 0-.68-.06-1.34-.14-2h3.38c.16.64.26 1.31.26 2s-.1 1.36-.26 2h-3.38z"/>
                        </svg>
                        Sitio Web
                    </a>
                    <a href="https://instagram.com/goalhero.app" class="social-link instagram">
                        <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor" style="margin-right: 8px;">
                            <path d="M7.8 2h8.4C19.4 2 22 4.6 22 7.8v8.4a5.8 5.8 0 0 1-5.8 5.8H7.8C4.6 22 2 19.4 2 16.2V7.8A5.8 5.8 0 0 1 7.8 2m-.2 2A3.6 3.6 0 0 0 4 7.6v8.8C4 18.39 5.61 20 7.6 20h8.8a3.6 3.6 0 0 0 3.6-3.6V7.6C20 5.61 18.39 4 16.4 4H7.6m9.65 1.5a1.25 1.25 0 0 1 1.25 1.25A1.25 1.25 0 0 1 17.25 8 1.25 1.25 0 0 1 16 6.75a1.25 1.25 0 0 1 1.65-1.25M12 7a5 5 0 0 1 5 5 5 5 0 0 1-5 5 5 5 0 0 1-5-5 5 5 0 0 1 5-5m0 2a3 3 0 0 0-3 3 3 3 0 0 0 3 3 3 3 0 0 0 3-3 3 3 0 0 0-3-3z"/>
                        </svg>
                        Instagram
                    </a>
                </div>
            </div>
            
            <div style="text-align: center; margin-top: 40px; padding-top: 30px; border-top: 1px solid #e5e7eb;">
                <p style="color: #6b7280; font-size: 16px;">
                    ¿Tienes preguntas? ¡Estamos aquí para ayudar! Responde a este email o contáctanos en 
                    <a href="mailto:info@goalhero.eu" style="color: #4CAF50;">info@goalhero.eu</a>
                </p>
            </div>
        </div>
        
        <div class="footer">
            <p><strong>Equipo GoalHero</strong></p>
            <p>Haciendo los sueños alcanzables, un gol a la vez.</p>
            <p style="margin-top: 20px; font-size: 14px; opacity: 0.8;">
                © 2025 GoalHero. Todos los derechos reservados.<br>
                <a href="#">Darse de baja</a> | <a href="#">Política de Privacidad</a> | <a href="#">Términos de Servicio</a>
            </p>
        </div>
    </div>
</body>
</html>
    `
}
