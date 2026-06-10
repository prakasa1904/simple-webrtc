# Simple WebRTC Project

## Overview

This repository implements a Clean Architecture-based Go backend using the following tools:

- **Fiber** for HTTP routing and server management.
- **GORM** for ORM-based database interactions.
- **Zap** for structured logging.
- **Validator** for input validation.

## Project Structure

The project adheres to the Clean Architecture principles:

```
internal/
├── app/
│   └── bootstra... (application wiring)
├── config/
│   └── config.go (configuration loading)
├── platform/
│   ├── database/ (GORM implementation)
│   ├── logger/ (Zap setup)
│   └── validator/ (input validation)
├── shared/
│   └── cross-feature utilities (middleware, pagination, etc.)
├── features/
│   ├── user/
│   │   ├── domain/ (business entities)
│   │   ├── usecase/ (application logic)
│   │   ├── repository/ (persistence layer)
│   │   └── delivery/http/ (Fiber handlers)
│   └── order/
│       └── ...
```

### KeyRules

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
