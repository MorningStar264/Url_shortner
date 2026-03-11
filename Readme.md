# Url_shortner
A fast, lightweight, and scalable URL shortening service written in Go. It converts long URLs into short, shareable links and handles redirections efficiently.

Features
  Shorten long URLs into compact Base62 IDs
  Fast redirection using Redis as backend storage
  RESTful API endpoints (/shorten and /[id])
  Custom alias support (optional)
  Easily deployable via Docker & Docker Compose
