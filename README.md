# DevDocs AI

**دستیار هوشمند و شخصی‌سازی‌شده برای جستجو و پرسش‌وپاسخ در مستندات فنی، توسعه‌یافته با Go، LangChainGo، Qdrant، Redis و Ollama.**

DevDocs AI یک برنامه مبتنی بر معماری **Retrieval-Augmented Generation (RAG)** است که به توسعه‌دهندگان اجازه می‌دهد مستندات فنی خود را بارگذاری کنند و با استفاده از زبان طبیعی درباره محتوای آن‌ها سؤال بپرسند.

به‌جای جستجوی دستی در میان فایل‌ها و مستندات مختلف، کاربران می‌توانند سؤال خود را مطرح کنند و پاسخ‌هایی مرتبط با محتوای اسناد شخصی خود دریافت کنند.

## معرفی پروژه

DevDocs AI با ترکیب جستجوی معنایی، مدل‌های زبانی محلی و حافظه مکالمه، یک دستیار هوشمند برای مدیریت و جستجو در دانش فنی شخصی ایجاد می‌کند.

### نحوه عملکرد

1. بارگذاری مستندات فنی مانند Markdown، TXT و PDF.
2. استخراج متن و تقسیم محتوا به بخش‌های کوچک‌تر (Chunking).
3. تولید بردارهای معنایی (Embeddings) با استفاده از Ollama.
4. ذخیره بردارها و اطلاعات اسناد در Qdrant.
5. دریافت سؤال کاربر از طریق Chat API.
6. بازیابی مرتبط‌ترین بخش‌های اسناد با Semantic Search.
7. تولید پاسخ متناسب با محتوای بازیابی‌شده توسط LLM.
8. ذخیره تاریخچه مکالمات در Redis.

## قابلیت‌ها

- **Document Ingestion:** بارگذاری و پردازش مستندات فنی.
- **Semantic Search:** جستجوی هوشمند براساس مفهوم و معنای متن.
- **RAG Pipeline:** تولید پاسخ براساس اطلاعات بازیابی‌شده از اسناد.
- **Conversational Memory:** مدیریت تاریخچه مکالمات با Redis.
- **Local LLM Integration:** اجرای مدل‌های زبانی به‌صورت محلی با Ollama.
- **Vector Storage:** ذخیره و بازیابی Embeddingها با Qdrant.
- **REST API:** ارائه قابلیت‌های پروژه از طریق API مبتنی بر Go و Fiber.
- **Modular Architecture:** معماری ماژولار با تفکیک Handlerها، Serviceها و Configuration.

## فناوری‌های استفاده‌شده

| فناوری | کاربرد |
|---|---|
| Go | توسعه Backend |
| Fiber | پیاده‌سازی REST API |
| LangChainGo | مدیریت فرایندهای LLM و RAG |
| Ollama | اجرای مدل‌های زبانی و Embedding |
| Qdrant | پایگاه داده برداری |
| Redis | مدیریت حافظه مکالمات و Cache |
| Docker | اجرای سرویس‌های زیرساختی |

## معماری پروژه

```text
                 Client / Postman
                        |
                        v
                 Go + Fiber API
                        |
                        v
                   LangChainGo
                        |
              +---------+---------+
              |                   |
              v                   v
        RAG Pipeline         Chat Memory
              |                   |
              v                   v
           Qdrant               Redis
              |
              v
            Ollama
       (Embeddings + LLM)
```

## ساختار پروژه

```text
devdocs-ai/
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── config/
│   │   └── config.go
│   │
│   ├── handlers/
│   │   ├── document.go
│   │   └── chat.go
│   │
│   ├── services/
│   │   ├── document.go
│   │   ├── chat.go
│   │   ├── qdrant.go
│   │   ├── redis.go
│   │   └── ollama.go
│   │
│   └── routes/
│       └── routes.go
│
├── documents/
├── .env.example
├── .gitignore
├── docker-compose.yml
├── go.mod
└── README.md
```

## مدل‌های هوش مصنوعی

در نسخه اولیه پروژه از مدل‌های محلی و ابزار Ollama استفاده می‌شود.

| مدل | کاربرد |
|---|---|
| `qwen3.5:4b` | تولید پاسخ و پردازش مکالمات |
| `embeddinggemma:300m` | تولید Embedding اسناد و پرسش‌ها |

**نکته:** برای جستجوی برداری، مدل Embedding هنگام ذخیره اسناد و هنگام پردازش پرسش‌ها باید یکسان باشد.

## API Endpoints

| متد | مسیر | توضیحات |
|---|---|---|
| POST | `/documents/upload` | بارگذاری مستندات |
| GET | `/documents` | دریافت فهرست مستندات |
| DELETE | `/documents/:id` | حذف یک سند |
| POST | `/chat` | ارسال سؤال به دستیار هوشمند |
| GET | `/chat/history/:session_id` | دریافت تاریخچه مکالمه |
| DELETE | `/chat/history/:session_id` | پاک کردن تاریخچه مکالمه |

## راه‌اندازی پروژه

### پیش‌نیازها

قبل از اجرای پروژه، ابزارهای زیر باید نصب باشند:

- Go
- Docker و Docker Compose
- Ollama

### ۱. دریافت سورس پروژه

```bash
git clone https://github.com/YOUR_USERNAME/devdocs-ai.git
cd devdocs-ai
```

### ۲. دریافت مدل‌های Ollama

```bash
ollama pull qwen3.5:4b
ollama pull embeddinggemma:300m
```

### ۳. تنظیم متغیرهای محیطی

فایل تنظیمات محیطی را ایجاد کنید:

```bash
cp .env.example .env
```

سپس تنظیمات مربوط به سرویس‌ها و مدل‌های مورد استفاده را در فایل `.env` وارد کنید.

### ۴. اجرای سرویس‌های زیرساختی

```bash
docker compose up -d
```

### ۵. اجرای برنامه

```bash
go mod tidy
go run ./cmd/server
```

## نقشه راه توسعه

این پروژه به‌صورت مرحله‌ای و با هدف یادگیری عملی توسعه داده می‌شود.

- [ ] **Day 01:** راه‌اندازی پروژه و مدیریت Configuration
- [ ] **Day 02:** اتصال Redis و پیاده‌سازی Chat History
- [ ] **Day 03:** اتصال Qdrant و مدیریت Vector Storage
- [ ] **Day 04:** بارگذاری اسناد، Chunking و تولید Embeddings
- [ ] **Day 05:** پیاده‌سازی RAG Pipeline
- [ ] **Day 06:** توسعه Chat API و Conversational Memory
- [ ] **Day 07:** یکپارچه‌سازی، تست و بهینه‌سازی نهایی

## اهداف آموزشی

این پروژه علاوه بر توسعه یک برنامه کاربردی، با هدف یادگیری عملی مفاهیم زیر طراحی شده است:

- معماری Retrieval-Augmented Generation (RAG)
- پایگاه‌های داده برداری (Vector Databases)
- جستجوی معنایی (Semantic Search)
- مدل‌های Embedding
- کامپوننت‌ها و Chainهای LangChainGo
- مدیریت حافظه مکالمات (Conversational Memory)
- اتصال Redis به Go
- اتصال Qdrant به Go
- طراحی معماری ماژولار برای Backend

## وضعیت پروژه

**وضعیت:** در حال توسعه (In Development)

پروژه DevDocs AI به‌صورت مرحله‌ای توسعه پیدا می‌کند و تمرکز اصلی آن بر درک صحیح مفاهیم، پیاده‌سازی اصولی اجزا و ایجاد یک سیستم RAG کاربردی است.

## مجوز استفاده

این پروژه تحت مجوز MIT منتشر خواهد شد.
