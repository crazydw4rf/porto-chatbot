package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"google.golang.org/genai"
)

const portfolioContext = `Berikut adalah portfolio singkat tentang Ucup:
Ucup adalah seorang Full Stack Developer dengan pengalaman lebih dari tiga tahun dalam mengembangkan aplikasi web modern. Ia memiliki keahlian yang solid di sisi front-end menggunakan React.js, Next.js, dan Vue.js, serta sisi back-end dengan Node.js, Express.js, dan Go. Dalam pengembangan sistem, ia terbiasa menerapkan arsitektur RESTful API maupun GraphQL, serta prinsip-prinsip clean architecture dan microservices. Ucup juga mahir dalam penggunaan berbagai database seperti PostgreSQL untuk transaksi yang kompleks, MongoDB untuk skema yang fleksibel, serta Redis sebagai caching layer untuk meningkatkan performa aplikasi.
Dalam perjalanan kariernya, Ucup telah menyelesaikan berbagai proyek yang memiliki dampak nyata. Salah satu proyek pentingnya adalah EduTrack, sebuah platform manajemen pembelajaran untuk lembaga kursus, di mana Ucup bertanggung jawab merancang dashboard dinamis untuk guru dan siswa, sistem penilaian otomatis, serta integrasi dengan payment gateway Midtrans. Proyek lainnya adalah PasarOnline, sebuah platform e-commerce lokal yang menghubungkan penjual UMKM dengan pelanggan. Di sini Ucup membangun sistem katalog produk, keranjang belanja, dan fitur live chat menggunakan WebSocket untuk interaksi real-time. Selain itu, ia juga mengembangkan MagicNotes, aplikasi pencatat pribadi berbasis web yang mendukung markdown, penyimpanan cloud (menggunakan AWS S3), dan mode kolaborasi real-time seperti Google Docs.
Tak hanya mengembangkan, Ucup juga aktif dalam proses deployment menggunakan Docker dan CI/CD pipeline dengan GitHub Actions. Ia terbiasa dengan environment staging dan production menggunakan layanan seperti Vercel, Heroku, dan DigitalOcean. Untuk kebutuhan testing, ia menggunakan Jest, Supertest, dan Postman untuk memastikan semua endpoint dan fitur bekerja dengan semestinya. Ia juga memiliki pengalaman dalam penulisan dokumentasi teknis menggunakan Swagger dan Notion, memudahkan tim lain memahami sistem yang dibangun.
Dengan semangat belajar yang tinggi dan pendekatan kerja yang terstruktur, Ucup terus mengeksplorasi teknologi baru seperti TypeScript, Prisma ORM, hingga arsitektur event-driven dengan Kafka. Ia percaya bahwa kualitas aplikasi tidak hanya ditentukan oleh fungsionalitas, tetapi juga oleh performa, maintainability, dan pengalaman pengguna yang menyenangkan.

Kamu adalah Chatbot yang dirancang untuk menjawab pertanyaan tentang portfolio Ucup. Kamu akan memberikan jawaban berdasarkan informasi yang telah diberikan di atas. Pastikan untuk menjawab dengan jelas dan sesuai konteks, tanpa menambahkan informasi lain yang tidak relevan.
Kamu hanya perlu menjawab pertanyaan berdasarkan informasi di atas dan jangan menambahkan informasi lain, cukup menjawab pertanyaan sesuai konteks yang diberikan.
`

type ChatRequest struct {
	Question string `json:"question"`
}

var (
	genaiClient  *genai.Client                = nil
	clientConfig *genai.GenerateContentConfig = &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(portfolioContext, genai.RoleUser),
	}
)

func getGenAIClient() (*genai.Client, error) {
	var err error
	if genaiClient == nil {
		ctx := context.Background()
		genaiClient, err = genai.NewClient(ctx, &genai.ClientConfig{
			APIKey:  os.Getenv("GEMINI_API_KEY"),
			Backend: genai.BackendGeminiAPI,
		})

		return genaiClient, err
	}

	return genaiClient, nil
}

func chatHandler(c *fiber.Ctx) error {
	chat := new(ChatRequest)

	err := c.BodyParser(chat)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	client, err := getGenAIClient()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create GenAI client")
	}

	result, err := client.Models.GenerateContent(
		c.Context(),
		"gemini-2.0-flash",
		genai.Text(chat.Question),
		clientConfig,
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate content")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"answer": result.Text()})
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendFile("./web/index.html")
	// })

	app.Static("/", "./web")

	app.Post("/chat", chatHandler)

	err := app.Listen(":3000")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
