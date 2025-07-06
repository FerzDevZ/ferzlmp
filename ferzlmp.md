// Project: FerzLmp - CLI-based Web Server Manager (Hybrid, Cross-Platform)
// Language: Go (Golang)
// Author: Ferdinand Dero
// Purpose: Build a lightweight, portable, CLI-only replacement for tools like Laragon, XAMPP, and FhyServee.
// Platform: Cross-platform (Windows and Linux)
// Type: CLI application only (no GUI, no REST API)

// ------------------------------------------------------------------------------------
// 🔥 OVERVIEW
// ------------------------------------------------------------------------------------
// FerzLmp is a CLI tool to manage a local PHP development server.
// It provides an all-in-one web environment (Apache, PHP, MySQL) that works without GUI.
// It supports both: 
//  1. Full bundled mode (comes with prepackaged binaries)
//  2. Installable mode (downloads Apache, PHP, MySQL via command)

// ------------------------------------------------------------------------------------
// 📦 TARGET USERS
// ------------------------------------------------------------------------------------
// - PHP developers (Laravel, WordPress, CodeIgniter, etc.)
// - Developers who prefer CLI over GUI
// - Users who need portable, offline-ready web servers
// - Developers on Windows or Linux who dislike setting up environments manually

// ------------------------------------------------------------------------------------
// ✅ CORE FEATURES
// ------------------------------------------------------------------------------------
// 1. Start and Stop Services:
//    - Command: `ferzlmp start`, `ferzlmp stop`
//    - Starts Apache and MySQL (from bundled binaries or installed paths)
//    - Checks for port conflicts (80, 3306)

// 2. Serve Projects:
//    - All web projects placed under `projects/` folder
//    - Auto map domain: `http://<project-name>.test`
//    - Configure Apache virtualhost and update local `hosts` file

// 3. Installer CLI:
//    - Command: `ferzlmp install php 8.2`, `install mysql 5.7`, etc.
//    - Downloads official binaries or portable versions
//    - Unzips them into `/modules/php/`, `/modules/mysql/`, etc.

// 4. Version Switcher:
//    - Command: `ferzlmp use php 8.1`, `use mysql 5.7`
//    - Sets active version globally or per-project

// 5. Project Generator:
//    - Command: `ferzlmp new laravel blog`, `new wordpress wp1`
//    - Downloads starter files and sets up vhost

// 6. Diagnostic Tool:
//    - Command: `ferzlmp doctor`
//    - Checks system ports, required binaries, permissions, virtualhost config, etc.

// 7. VirtualHost Manager:
//    - Command: `ferzlmp vhost add blog.test ./projects/blog`
//    - Edits Apache config and local `hosts` file automatically
//    - Removes vhost with `ferzlmp vhost remove blog.test`

// 8. Configuration via YAML:
//    - File: `config/ferzlmp.yaml`
//    - Stores: active PHP path, MySQL path, Apache path, project root, port config, etc.
//    - Supports default config + user override

// 9. Log and CLI User Experience:
//    - Color-coded terminal logs using `fatih/color`
//    - Friendly messages with icons, durations, success/failure states

// 10. CLI-Only, No GUI/API:
//    - Do not implement any graphical UI or web-based panel
//    - No HTTP server or REST API should be implemented

// ------------------------------------------------------------------------------------
// 🧱 PROJECT STRUCTURE
// ------------------------------------------------------------------------------------
// ferzlmp/
// ├── main.go
// ├── go.mod
// ├── /cmd/                 → Cobra CLI commands
// │   ├── root.go
// │   ├── start.go
// │   ├── stop.go
// │   ├── install.go
// │   ├── use.go
// │   ├── new.go
// │   ├── vhost.go
// │   └── doctor.go
// ├── /internal/            → Internal helper packages
// │   ├── config/           → YAML config loader
// │   ├── services/         → Start/stop logic for Apache, MySQL
// │   ├── download/         → Binary download + unzip
// │   ├── vhost/            → Virtualhost writer
// │   └── utils.go
// ├── /projects/            → Folder where user’s PHP apps live
// ├── /modules/             → PHP, MySQL, Apache installations
// ├── /config/ferzlmp.yaml  → Config file
// └── README.md

// ------------------------------------------------------------------------------------
// 🛠️ INSTALLATION MODES
// ------------------------------------------------------------------------------------
// For End-Users (Non-programmers):
// 1. Download ZIP release from GitHub Releases (includes ferzlmp binary + apache/php/mysql folders)
// 2. Extract to any folder (e.g. D:\FerzLmp or ~/ferzlmp)
// 3. Open Terminal/CMD and run:
//    - `ferzlmp start` → Starts Apache and MySQL
//    - `ferzlmp new laravel blog` → Creates Laravel app
//    - `ferzlmp doctor` → Diagnoses environment
// 4. Access apps at: http://blog.test

// For Developers:
// - Option 1 (Recommended):
//     `go install github.com/ferzdev/ferzlmp@latest`
// - Option 2 (Build manually):
//     `git clone https://github.com/ferzdev/ferzlmp.git`
//     `cd ferzlmp && go mod tidy && go run main.go`

// System Requirements:
// - Golang 1.18+ (for devs)
// - Windows or Linux (tested on Ubuntu 20.04+ and Windows 10/11)
// - Admin/root privileges required to modify hosts file and open port 80

// ------------------------------------------------------------------------------------
// 🤖 RULES FOR GITHUB COPILOT
// ------------------------------------------------------------------------------------
// - Write modular Go code, readable and maintainable
// - Use Cobra for CLI command parsing (`github.com/spf13/cobra`)
// - Use Viper or yaml.v2 for config loading (`github.com/spf13/viper`)
// - Use `os/exec` to start/stop Apache/MySQL binaries
// - Use `runtime.GOOS` to detect platform
// - Use `fatih/color` for colorful terminal logs
// - Split commands into `/cmd/*.go` and helpers into `/internal`
// - Use `filepath.Join()` instead of hardcoded paths
// - Add CLI help text for every command
// - Do not use any GUI libraries or REST API
// - Always validate user input and show clear error if invalid
// - Make all commands idempotent and safe to re-run
// - Handle Windows and Linux differences in paths and processes
// - If port 80/3306 already used, suggest alternative or show message

// ------------------------------------------------------------------------------------
// ✅ FIRST TASKS FOR COPILOT
// ------------------------------------------------------------------------------------
// 1. Scaffold CLI using Cobra with root command
// 2. Implement `ferzlmp start` to run Apache and MySQL
// 3. Implement config loader for `ferzlmp.yaml`
// 4. Implement `install` and `use` commands
// 5. Implement virtualhost creation in Apache
// 6. Check port usage before starting services (doctor)
// 7. Create sample Laravel & WordPress generators