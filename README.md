# FerzLmp

FerzLmp is a CLI-based web server manager for PHP development (Apache, PHP, MySQL) for Windows and Linux.

## Features
- Start/stop Apache & MySQL
- Install PHP/MySQL/Apache
- Switch PHP/MySQL version (global/per-project)
- Project generator (Laravel/WordPress)
- Virtualhost manager
- Diagnostic tool

## Quick Start
```sh
ferzlmp install php 8.2
ferzlmp install mysql 5.7
ferzlmp install apache 2.4
ferzlmp new laravel blog
ferzlmp vhost add blog.test ./projects/blog
ferzlmp start
ferzlmp doctor
```

## Requirements
- Go 1.18+
- curl, unzip, tar, composer (for project generator)
- Admin/root privileges for hosts file

## Project Structure
See `ferzlmp.md` for full details.
# ferzlmp
# ferzlmp
