# simple-webrtc - Clean Architecture Backend

This repository implements a Clean Architecture-based Go backend using the following tools:

- **Fiber** for HTTP routing and server management.
- **GORM** for ORM-based database interactions.
- **Zap** for structured logging.
- **Validator** for input validation.

## Project Structure

The project adheres to the Clean Architecture principles:

```
cmd/api/
  └── main.go (executable entrypoint)

internal/app/
  └── bootstrap.go (application wiring)

internal/config/
  └── config.go (configuration loading)

internal/platform/
  ├── database/gorm.go (database layer)
  ├── logger/zap.go (logging setup)
  └── validator/validator.go (input validation)

internal/shared/
  ├── errors/
  ├── response/
  └── middleware/

internal/features/user/
  ├── domain/user.go (business entities)
  ├── usecase/user_usecase.go (application logic)
  ├── repository/gorm_repository.go (persistence)
  └── delivery/http/handler.go (Fiber handlers)

internal/features/order/
  └── ... (similar structure)

migrations/
  └── database migrations

docs/
  └── technical documentation

scripts/
  └── developer/deployment scripts
```

### Key Rules

- **Framework dependence should point inward** (`delivery/http` → `usecase` → `domain`).
- **DO NOT** import Fiber, GORM, or Zap in `domain` or `usecase` layers.
- **Validation** is enforced in `delivery/http` with DTOs.
- **Logs** are handled in `platform/logger/` with Zap.
- **Database models** live in `platform/database/`.

## Setup

```sh
git clone https://github.com/prakasa1904/simple-webrtc.git
cd simple-webrtc
go mod tidy
```

## Documentation

- **README.md**: Project overview and setup instructions.
- **AGENTS.md**: Clean Architecture guidelines and tool requirements.

## Testing

- Unit test usecases.
- Mock repository interfaces.
- Test validation rules where useful.
- Test repository behavior with integration tests when database access is needed.
- Do NOT require Fiber for usecase tests.

## Notes

- This repository follows the same Clean Architecture rules as the [golang-webapp-boilerplate](https://github.com/devetek/golang-webapp-boilerplate).
- Avoid large global folders like `internal/usecase` or `internal/repository` for mature systems.
- Use feature-specific folders like `internal/features/user` for better organization.

---

### Commit Guidelines

All commits must follow the Conventional Commits format: `feat:`, `fix:`, or `chore:` followed by a description. This ensures semantic versioning and clear project history.
