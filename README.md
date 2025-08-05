# GoalHero Emailer API

A Golang API service deployed on Vercel for sending welcome emails to beta users.

## Features

- 🚀 **Beta Registration**: POST endpoint `/api/beta-register` to register users for beta
- 📧 **Beautiful Emails**: HTML/CSS styled welcome emails with logo and branding
- 🔗 **Social Media Links**: Integrated social media links in emails
- ⚡ **Vercel Deployment**: Optimized for serverless deployment on Vercel

## Setup

1. **Clone the repository**:
   ```bash
   git clone <your-repo-url>
   cd goalhero-emailer
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Environment Variables**:
   - Copy `.env.example` to `.env.local`
   - Choose one email method:
   
   **Option A: SMTP (Recommended - Free)**
   - Use your Gmail, Yahoo, or Outlook account
   - Set `SMTP_HOST`, `SMTP_PORT`, `SMTP_USER`, `SMTP_PASS`, `FROM_EMAIL`, `FROM_NAME`
   - For Gmail: Enable 2FA and create an [App Password](https://myaccount.google.com/apppasswords)
   
   **Option B: SendGrid API**
   - Get your SendGrid API key from [SendGrid Dashboard](https://app.sendgrid.com/settings/api_keys)
   - Set `SENDGRID_API_KEY`
   - Use the original `beta-register.go` file

4. **For Vercel deployment**:
   - Set your chosen environment variables in Vercel project settings
   - The `vercel.json` configuration is already set up

## API Endpoints

### POST /api/beta-register

Registers a user for the beta program and sends a welcome email.

**Request Body**:
```json
{
  "email": "user@example.com"
}
```

**Response**:
```json
{
  "success": true,
  "message": "Welcome email sent successfully!"
}
```

**Error Response**:
```json
{
  "success": false,
  "message": "Error description"
}
```

## Local Development

To test locally, you can use tools like curl:

```bash
curl -X POST http://localhost:3000/api/beta-register \
  -H "Content-Type: application/json" \
  -d '{"email": "test@example.com"}'
```

## Deployment

1. Connect your repository to Vercel
2. Set the `SENDGRID_API_KEY` environment variable in Vercel dashboard
3. Deploy - Vercel will automatically detect the Go functions

## Email Template

The welcome email includes:
- ✨ Beautiful responsive HTML/CSS design
- 🎨 GoalHero branding and logo
- 🎉 Congratulations message for beta registration
- 📋 Feature preview list
- 🔗 Social media links (Twitter, LinkedIn, Instagram, Facebook)
- 📞 Support contact information
- 📱 Mobile-responsive design

## Technologies Used

- **Go 1.21**: Backend API
- **SendGrid**: Email delivery service
- **Vercel**: Serverless deployment platform
- **HTML/CSS**: Email template styling