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
	from := mail.NewEmail("GoalHero Team", "noreply@goalhero.com")
	to := mail.NewEmail("", email)
	subject := "🎉 Welcome to GoalHero Beta!"

	htmlContent := getWelcomeEmailHTML()
	
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	if response.StatusCode >= 400 {
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
            border-radius: 12px;
            overflow: hidden;
            box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
        }
        
        .header {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            padding: 40px 30px;
            text-align: center;
        }
        
        .logo {
            width: 80px;
            height: 80px;
            background-color: #ffffff;
            border-radius: 50%;
            margin: 0 auto 20px;
            display: flex;
            align-items: center;
            justify-content: center;
            font-size: 36px;
            font-weight: bold;
            color: #667eea;
        }
        
        .header h1 {
            color: #ffffff;
            font-size: 28px;
            font-weight: 700;
            margin-bottom: 10px;
        }
        
        .header p {
            color: #e2e8f0;
            font-size: 16px;
        }
        
        .content {
            padding: 40px 30px;
        }
        
        .welcome-message {
            text-align: center;
            margin-bottom: 30px;
        }
        
        .welcome-message h2 {
            font-size: 24px;
            color: #2d3748;
            margin-bottom: 15px;
        }
        
        .welcome-message p {
            font-size: 16px;
            color: #718096;
            line-height: 1.8;
        }
        
        .features {
            background-color: #f7fafc;
            border-radius: 8px;
            padding: 25px;
            margin: 30px 0;
        }
        
        .features h3 {
            font-size: 18px;
            color: #2d3748;
            margin-bottom: 15px;
            text-align: center;
        }
        
        .feature-list {
            list-style: none;
        }
        
        .feature-list li {
            padding: 8px 0;
            color: #4a5568;
            display: flex;
            align-items: center;
        }
        
        .feature-list li:before {
            content: "✨";
            margin-right: 10px;
            font-size: 16px;
        }
        
        .cta-section {
            text-align: center;
            margin: 30px 0;
        }
        
        .cta-button {
            display: inline-block;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: #ffffff;
            text-decoration: none;
            padding: 15px 30px;
            border-radius: 8px;
            font-weight: 600;
            font-size: 16px;
            transition: transform 0.2s ease;
        }
        
        .cta-button:hover {
            transform: translateY(-2px);
        }
        
        .social-section {
            background-color: #f7fafc;
            border-radius: 8px;
            padding: 25px;
            text-align: center;
            margin: 30px 0;
        }
        
        .social-section h3 {
            font-size: 18px;
            color: #2d3748;
            margin-bottom: 15px;
        }
        
        .social-links {
            display: flex;
            justify-content: center;
            gap: 20px;
            flex-wrap: wrap;
        }
        
        .social-link {
            display: inline-block;
            color: #667eea;
            text-decoration: none;
            font-weight: 500;
            padding: 8px 16px;
            border: 2px solid #667eea;
            border-radius: 25px;
            transition: all 0.3s ease;
        }
        
        .social-link:hover {
            background-color: #667eea;
            color: #ffffff;
        }
        
        .footer {
            background-color: #2d3748;
            color: #cbd5e0;
            padding: 30px;
            text-align: center;
        }
        
        .footer p {
            margin-bottom: 10px;
            font-size: 14px;
        }
        
        .footer a {
            color: #667eea;
            text-decoration: none;
        }
        
        @media (max-width: 600px) {
            .container {
                margin: 10px;
                border-radius: 8px;
            }
            
            .header, .content {
                padding: 30px 20px;
            }
            
            .header h1 {
                font-size: 24px;
            }
            
            .welcome-message h2 {
                font-size: 20px;
            }
            
            .social-links {
                flex-direction: column;
                align-items: center;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <div class="logo">GH</div>
            <h1>Welcome to GoalHero!</h1>
            <p>You're now part of our exclusive beta community</p>
        </div>
        
        <div class="content">
            <div class="welcome-message">
                <h2>🎉 Congratulations!</h2>
                <p>Thank you for joining the GoalHero beta program! You're among the first to experience our revolutionary goal-tracking platform that transforms how you achieve your dreams.</p>
            </div>
            
            <div class="features">
                <h3>What's Coming Your Way</h3>
                <ul class="feature-list">
                    <li>Intelligent goal tracking with AI-powered insights</li>
                    <li>Progress visualization and milestone celebrations</li>
                    <li>Community challenges and accountability partners</li>
                    <li>Personalized coaching and motivation</li>
                    <li>Advanced analytics and performance metrics</li>
                </ul>
            </div>
            
            <div class="cta-section">
                <a href="#" class="cta-button">Get Started Now</a>
            </div>
            
            <div class="social-section">
                <h3>Stay Connected</h3>
                <div class="social-links">
                    <a href="https://twitter.com/goalhero" class="social-link">Twitter</a>
                    <a href="https://linkedin.com/company/goalhero" class="social-link">LinkedIn</a>
                    <a href="https://instagram.com/goalheroapp" class="social-link">Instagram</a>
                    <a href="https://facebook.com/goalheroapp" class="social-link">Facebook</a>
                </div>
            </div>
            
            <div style="text-align: center; margin-top: 30px; padding-top: 20px; border-top: 1px solid #e2e8f0;">
                <p style="color: #718096; font-size: 14px;">
                    Have questions? We're here to help! Reply to this email or contact us at 
                    <a href="mailto:support@goalhero.com" style="color: #667eea;">support@goalhero.com</a>
                </p>
            </div>
        </div>
        
        <div class="footer">
            <p><strong>GoalHero Team</strong></p>
            <p>Making dreams achievable, one goal at a time.</p>
            <p style="margin-top: 15px; font-size: 12px; opacity: 0.8;">
                © 2025 GoalHero. All rights reserved.<br>
                <a href="#">Unsubscribe</a> | <a href="#">Privacy Policy</a> | <a href="#">Terms of Service</a>
            </p>
        </div>
    </div>
</body>
</html>
    `
}