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
	Email string `json:"email"`
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

	if err := sendWelcomeEmail(req.Email); err != nil {
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

func sendWelcomeEmail(email string) error {
	from := mail.NewEmail("GoalHero Team", "info@goalhero.eu")
	to := mail.NewEmail("", email)
	subject := "🎉⚽ Welcome to GoalHero Beta!"

	htmlContent := getWelcomeEmailHTML()

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
    <title>Welcome to GoalHero Beta</title>
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
            background: linear-gradient(135deg, #4CAF50 0%, #2E7D32 100%);
            color: #ffffff;
            text-decoration: none;
            padding: 18px 40px;
            border-radius: 12px;
            font-weight: 600;
            font-size: 18px;
            transition: all 0.3s ease;
            box-shadow: 0 4px 15px rgba(76, 175, 80, 0.3);
        }
        
        .cta-button:hover {
            transform: translateY(-3px);
            box-shadow: 0 8px 25px rgba(76, 175, 80, 0.4);
        }
        
        .social-section {
            background: linear-gradient(135deg, #f8f8f8 0%, #f0f0f0 100%);
            border-radius: 16px;
            padding: 30px;
            text-align: center;
            margin: 40px 0;
            border: 1px solid #e0e0e0;
        }
        
        .social-section h3 {
            font-size: 22px;
            color: #1a1a1a;
            margin-bottom: 25px;
            font-weight: 600;
        }
        
        .social-links {
            display: flex;
            justify-content: center;
            gap: 15px;
            flex-wrap: wrap;
        }
        
        .social-link {
            display: inline-block;
            color: #4CAF50;
            text-decoration: none;
            font-weight: 500;
            padding: 12px 20px;
            border: 2px solid #4CAF50;
            border-radius: 30px;
            transition: all 0.3s ease;
            font-size: 16px;
        }
        
        .social-link:hover {
            background-color: #4CAF50;
            color: #ffffff;
            transform: translateY(-2px);
            box-shadow: 0 4px 15px rgba(76, 175, 80, 0.3);
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
            </div>
            
            <div class="features">
                <h3>What's Coming Your Way</h3>
                <ul class="feature-list">
                    <li>Post games and receive competitive bids from goalkeepers</li>
                    <li>Browse verified goalkeeper profiles with ratings & reviews</li>
                    <li>Secure payment system with guaranteed show-up protection</li>
                    <li>24/7 emergency backup keeper network</li>
                    <li>Match analytics and keeper performance tracking</li>
                </ul>
            </div>
            
            <div class="cta-section">
                <a href="https://www.goalhero.eu" class="cta-button">Get Started Now</a>
            </div>
            
            <div class="social-section">
                <h3>Stay Connected</h3>
                <div class="social-links">
                    <a href="https://www.goalhero.eu" class="social-link">🌐 Website</a>
                    <a href="https://instagram.com/goalhero.app" class="social-link">📸 Instagram</a>
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
