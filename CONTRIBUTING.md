# Contributing to ClassWork 🚀

First off, thank you for considering contributing to **ClassWork**! 🎉 Open source projects thrive because of contributors like you.

Whether you're fixing a bug, proposing a new feature, improving documentation, or refining UI/UX, all contributions are welcome and greatly appreciated.

---

## 📜 Table of Contents

1. [Code of Conduct](#-code-of-conduct)
2. [Project Overview & Architecture](#-project-overview--architecture)
3. [Prerequisites](#-prerequisites)
4. [Setting Up Your Local Development Environment](#-setting-up-your-local-development-environment)
   - [Clone the Repository](#1-clone-the-repository)
   - [Backend Setup (Go)](#2-backend-setup-go)
   - [Frontend Setup (Angular)](#3-frontend-setup-angular)
5. [Development Workflow](#-development-workflow)
   - [Branch Naming](#branch-naming-conventions)
   - [Commit Messages](#commit-message-conventions)
6. [Code Style & Best Practices](#-code-style--best-practices)
   - [Backend (Go)](#backend-guidelines)
   - [Frontend (Angular / TypeScript)](#frontend-guidelines)
7. [Testing](#-testing)
8. [Submitting a Pull Request (PR)](#-submitting-a-pull-request-pr)
9. [Reporting Bugs & Requesting Features](#-reporting-bugs--requesting-features)
10. [Need Help?](#-need-help)

---

## 🤝 Code of Conduct

This project and everyone participating in it is governed by our [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior to the project maintainers.

---

## 🏛️ Project Overview & Architecture

ClassWork is structured as a monorepo consisting of:

- **`backend/`**: A RESTful API written in **Go** implementing **Hexagonal Architecture (Ports & Adapters)**.
  - `internal/domain`: Core domain models and repository contracts (no external dependencies).
  - `internal/ports`: Inbound and outbound application port interfaces.
  - `internal/application`: Business logic and orchestration services (Auth, Homework, AI).
  - `internal/adapters`: Concrete drivers (HTTP with Chi router, PostgreSQL persistence, Gemini AI client).
  - `pkg/openapi`: OpenAPI 3.0 specification.
- **`frontend/`**: A responsive web application built with **Angular 21** and **Angular Material**, supporting teacher authentication, homework management, and AI-powered assignment generation.

---

## 🛠️ Prerequisites

Before you begin, ensure you have the following tools installed:

- **Git** 2.x+
- **Go** 1.21+ (1.25 recommended)
- **Node.js** 20+ & **npm** 10+
- **PostgreSQL** 12+ (or Docker for running PostgreSQL)
- *(Optional)* **Google Gemini API Key** for testing AI assignment generation features

---

## 💻 Setting Up Your Local Development Environment

### 1. Clone the Repository

```bash
git clone https://github.com/MrPerfect8699/class-work.git
cd class-work
```

### 2. Backend Setup (Go)

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Create a local environment configuration:
   ```bash
   cp .env.example .env   # Or create .env with your credentials
   ```

3. Configure your `.env` file:
   ```dotenv
   # Server
   SERVER_PORT=8080
   SERVER_HOST=0.0.0.0

   # Database (PostgreSQL)
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=classwork_db
   DB_SSLMODE=disable

   # JWT Secret
   JWT_SECRET=your_jwt_secret_key_here
   JWT_EXPIRATION_HOURS=24

   # AI Integration (Optional)
   GEMINI_API_KEY=your_gemini_api_key_here
   GEMINI_MODEL=gemini-2.5-flash
   ```

4. Download dependencies and run the server:
   ```bash
   go mod download
   go run main.go
   ```

   The backend will connect to PostgreSQL, automatically initialize the required schema, and listen on `http://localhost:8080`.

### 3. Frontend Setup (Angular)

1. In a new terminal, navigate to the frontend app:
   ```bash
   cd frontend/classWork-ui
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the Angular development server:
   ```bash
   npm start
   # or
   npx ng serve
   ```

4. Open your browser at `http://localhost:4200` to interact with the UI.

---

## 🔄 Development Workflow

### Branch Naming Conventions

Create feature branches off the main branch (`migration-to-postgresql` or `main`) with descriptive names:

- `feat/<feature-name>`: New feature or enhancement (e.g., `feat/rubric-generator`)
- `fix/<bug-name>`: Bug fix (e.g., `fix/jwt-token-expiration`)
- `docs/<doc-name>`: Documentation changes (e.g., `docs/api-guide`)
- `refactor/<refactor-name>`: Code refactoring without behavioral change
- `test/<test-name>`: Adding or modifying test suites
- `chore/<task-name>`: Tooling, dependencies, or repository maintenance

### Commit Message Conventions

We recommend following the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>(<scope>): <short summary>

[optional body]

[optional footer(s)]
```

**Types:**
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code (white-space, formatting)
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `perf`: A code change that improves performance
- `test`: Adding missing tests or correcting existing tests
- `chore`: Changes to the build process or auxiliary tools

*Example:* `feat(ai): add temperature configuration for assignment generation`

---

## 📐 Code Style & Best Practices

### Backend Guidelines

- **Adhere to Hexagonal Boundaries**: Never import `adapters` into the `domain` layer. Domain entities must remain pure.
- **Go Formatting**: Format all Go code before committing using `gofmt` or `goimports`:
  ```bash
  go fmt ./...
  ```
- **Static Analysis**: Run `go vet` to catch potential bugs:
  ```bash
  go vet ./...
  ```
- **Error Handling**: Return clear and idiomatic Go errors. Do not swallow errors silently.

### Frontend Guidelines

- **Code Formatting**: Use Prettier to keep code cleanly styled:
  ```bash
  npx prettier --write "src/**/*.{ts,html,scss,css}"
  ```
- **Component Design**: Keep Angular components focused, modular, and leverage Angular Material design guidelines.
- **State & Observables**: Unsubscribe or manage RxJS subscriptions appropriately to avoid memory leaks.

---

## 🧪 Testing

Always ensure that tests pass before pushing code.

### Backend Tests
Run the Go test suite:
```bash
cd backend
go test -v ./...
```

### Frontend Tests
Run the Angular test suite:
```bash
cd frontend/classWork-ui
npm test
```

---

## 📬 Submitting a Pull Request (PR)

1. **Fork & Branch**: Fork the repo and create your branch from the latest development branch.
2. **Make Changes**: Implement your changes adhering to project coding standards and architecture.
3. **Test**: Ensure all unit and integration tests pass.
4. **Commit & Push**: Push your branch to your fork.
5. **Open a PR**:
   - Provide a concise title and thorough description of what the PR solves.
   - Reference any related issue numbers (e.g., `Fixes #42`).
   - Include before/after screenshots or terminal outputs for visual or API changes.
6. **Code Review**: Engage constructively during code review and address requested changes.

---

## 🐛 Reporting Bugs & Requesting Features

### Found a Bug?
- Open an issue on GitHub.
- Use a clear title and provide step-by-step reproduction steps, expected behavior, actual behavior, and relevant logs/system environment information.

### Have a Feature Idea?
- Open an issue titled `[Feature Request] <Brief Summary>`.
- Explain the motivation, use case, and proposed implementation details.

---

## 💡 Need Help?

If you have questions or get stuck, feel free to open a discussion or reach out to the maintainers through issues or at [support@classwork.local](mailto:support@classwork.local).

Happy coding and thank you for contributing to ClassWork! 🎓✨
