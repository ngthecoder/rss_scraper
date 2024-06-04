# RSS Scraper API
This project sets up an HTTP server using the Chi router in Go, providing various endpoints for handling user and feed data. It connects to a PostgreSQL database and uses middleware for authentication. The API allows users to manage feeds, follow them, and fetch posts from the feeds. It also includes an RSS scraper that periodically scrapes feeds and updates the database.
## Prerequisites
- Go programming language (version 1.16 or later)
- PostgreSQL
- Git (for cloning the repository)
- A .env file with the following variables
```
PORT=your_port_number
DB_URL=your_database_url
```
## Files Overview
1. main.go: Initializes the server, sets up routes, and configures the database connection.
2. models.go: Defines the data models used in the application.
3. rss.go: Contains the logic for processing and storing RSS feed data.
4. scraper.go: Manages the periodic scraping of RSS feeds.
5. json.go: Contains helper functions for JSON encoding and decoding.
6. middleware_auth.go: Provides middleware for API key authentication.
7. handler_readiness.go: Provides a health check endpoint.
8. handler_err.go: Handles error simulation endpoints.
9. handler_user.go: Manages user creation and retrieval endpoints.
10. handler_feed.go: Handles feed creation and retrieval endpoints.
11. handler_feed_follows.go: Manages endpoints for following and unfollowing feeds.
## Getting Started
1. Clone the repository:
```
git clone https://github.com/yourusername/rss-scraper.git
```
2. Navigate to the project directory:
```
cd rss-scraper
```
3. Install dependencies
```
go mod tidy
```
4. Build the project:
```
go build
```
5. Run the application:
```
./rss-scraper
```
The server will start running on the specified port (e.g., http://localhost:8080).
## Endpoints
### Authentication
When authentication is required, users must add an Authorization header with the value `ApiKey {your_api_key_here}`.
### Health Check
#### GET /v1/healthz
- Description: Checks the readiness of the server.
- Response: 200 OK
### Error Simulation
#### GET /v1/err
- Description: Triggers an error for testing purposes.
- Response: 500 Internal Server Error
### User Management
#### POST /v1/users
- Description: Creates a new user.
- Request Body: JSON object containing user data.
```json
{
    "name": "{name}"
}
```
- Response: 201 Created
#### GET /v1/users
- Description: Retrieves the current user.
- Authentication: Required.
- Response: 200 OK with user data.
### Feed Management
#### POST /v1/feeds
- Description: Creates a new feed.
- Request Body: JSON object containing feed data.
```json
{
    "name": "{feed name}",
    "url": "{feed URL}"
}
```
- Authentication: Required.
- Response: 201 Created
#### GET /v1/feeds
- Description: Retrieves all feeds.
- Response: 200 OK with a list of feeds.
### Post Management
#### GET /v1/posts
- Description: Retrieves posts for the current user.
- Authentication: Required.
- Response: 200 OK with a list of posts.
### Feed Follows Management
#### POST /v1/feed_follows
- Description: Follows a feed.
- Request Body: JSON object containing feed follow data.
```json
{
    "feed_id": "{feed id}"
}
```
- Authentication: Required.
- Response: 201 Created
#### GET /v1/feed_follows
- Description: Retrieves the feeds that the current user follows.
- Authentication: Required.
- Response: 200 OK with a list of followed feeds.
#### DELETE /v1/feed_follows/{feedFollowID}
- Description: Unfollows a feed.
- URL Parameter: feedFollowID - ID of the feed follow to delete.
- Authentication: Required.
- Response: 204 No Content
## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.