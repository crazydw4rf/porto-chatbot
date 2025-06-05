# porto-chatbot

## How to Build & Run the Application

1. **Clone this repository**
   ```bash
   git clone https://github.com/crazydw4rf/porto-chatbot.git
   cd porto-chatbot
   ```

2. **Install dependencies**
   Make sure you have Go installed (minimum version 1.20).
   Then run:
   ```bash
   go mod download
   ```

3. **Configure .env file**
   Copy `.env.example` to `.env` in the root folder:
   ```bash
   cp .env.example .env
   ```
   Then, open the `.env` file and enter a valid Gemini API key for the `GEMINI_API_KEY` variable.

4. **Build the application**
   ```bash
   go build -o porto-chatbot ./cmd/porto-chatbot/main.go
   ```

5. **Run the application**
   ```bash
   ./porto-chatbot
   ```
   By default, the application will run on port 3000.

## Endpoint

- `POST /chat`
  Example request payload:
  ```json
  {
    "prompt": "What is Ucup's experience with Docker?"
  }
  ```
  Example Response:
  ```json
  {
    "response": "Ucup is active in deployment processes using Docker and CI/CD pipelines with GitHub Actions..."
  }
  ```

## Deploy to Vercel Serverless Function

This project is configured to be deployed to Vercel as a serverless function. Follow these steps:

### 1. Set Environment Variables in Vercel

After creating a project in Vercel:

1. Open your Vercel project dashboard
2. Go to **Settings** → **Environment Variables**
3. Add the following variables:
   - `GEMINI_API_KEY`: Your Gemini API key
   - `CORS_DOMAINS`: Allowed domains (example: `https://yourapp.vercel.app` or `*` for all)

### 2. Deploy

There are two ways to deploy:

**Method 1: Through GitHub (Recommended)**
1. Push the project to a GitHub repository
2. Import the repository in Vercel dashboard
3. Vercel will automatically build and deploy

**Method 2: Through Vercel CLI**
```bash
# Install Vercel CLI
npm i -g vercel

# Deploy
vercel

# Follow the prompts
```

### 3. Testing

After successful deployment, you can test the endpoint:
```bash
curl -X POST https://your-app.vercel.app/chat \
  -H "Content-Type: application/json" \
  -d '{"prompt": "What is Ucup'\''s experience with Docker?"}'
```

## Notes

- The `web/` folder contains frontend files (e.g., index.html).
- If you want to change the port, edit the following part in `main.go`:
  ```go
  err := app.Listen(":3000")
  ```
  Replace `3000` with your desired port.

## Modifying Portfolio Instructions

If you want to change the portfolio context used by the chatbot, modify the `UcupPortfolio` constant value in the `instruction/portfolio.go` file.

Example:
```go
const UcupPortfolio = `Here is a brief portfolio about Ucup:
Ucup is a Backend Developer experienced in building APIs with Go and PostgreSQL...`
```
Make sure to adjust the content according to your needs.
