# URL Shortener

##  Project Requirement: URL Shortener Service

We're building a simple URL Shortener backend in Golang.
The goal is to take long URLs and convert them into short, shareable links, with basic stats support.

---

###  Core Features

#### 1. Shorten a URL

* **Endpoint**: `POST /shorten`
* **Request Body**:

  ```json
  {
    "url": "https://www.youtube.com/watch?v=abcd"
  }
  ```
* **Response**:

  ```json
  {
    "short_url": "http://localhost:8080/abc123"
  }
  ```
* If the URL was already shortened before, return the same short code (avoid duplicates).

---

#### 2. Redirect to the Original URL

* **Endpoint**: `GET /<short_code>`
* Should look up the original URL for the given code and redirect using `302 Found`.

---

#### 3. Domain Metrics (Basic Stats)

* **Endpoint**: `GET /metrics`
* Response should include the top 3 domains that have been shortened the most, like:

  ```json
  [
    {"domain": "youtube.com", "count": 10},
    {"domain": "google.com", "count": 6},
    {"domain": "wikipedia.org", "count": 3}
  ]
  ```

---

###  Tech Requirements

* Use Go's built-in `net/http` package (no frameworks)
* Store data in memory using `map[string]string`
* Use a `sync.Mutex` to handle concurrent access safely
* For the short code, generate a random 6-character string
* Use `logrus` for logging

---

###  Optional Enhancements (if time allows)

* Validate that the input URL is valid and reachable
* Add unit tests for handlers
* Dockerize the service

---
