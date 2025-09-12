Go Web Server without any third party dependencies

## Overview

This folder contains a minimal HTTP server built with Go's standard library (`net/http`). It exposes simple in-memory CRUD-style endpoints for a `User` resource defined in `main.go`.

Key notes:
- Runs on http://localhost:8080
- In-memory storage (resets when the process restarts)
- Endpoints use Go's ServeMux patterns (Go 1.22+ style like `GET /users/{id}`)

## Requirements

- Go 1.22+ (module declares Go 1.25)
- Windows PowerShell or any shell that can run `go` and send HTTP requests

## Run the server

From the repository root:

```pwsh
# Option A: run by package path
go run ./Web_Server

# Option B: run by file path
go run Web_Server/main.go
```

You should see:

```
Starting server on :8080
```

Visit http://localhost:8080/ to verify it responds with "Welcome to the root!".

## API Endpoints

Base URL: http://localhost:8080

1) GET `/` — Health/root check
- Response: 200 OK, text/plain

2) POST `/users` — Create a user
- Body (JSON):
	- Required fields: `id` (number), `name` (string), `age` (number)
- Response: 204 No Content
	- Note: The current implementation writes a small body with 204; many clients ignore this by spec.

3) GET `/users/{id}` — Get a user by ID
- Path params: `id` (number)
- Response: 200 OK, application/json
	- JSON keys are mixed-cased due to tags (e.g., `ID`, `name`, `Age`). Example:
		```json
		{ "ID": 1, "name": "Alice", "Age": 30 }
		```
- Errors: 400 (bad id), 404 (not found)

4) DELETE `/users/{id}` — Delete a user by ID
- Response: 204 No Content
- Errors: 400 (bad id), 404 (not found)

Important implementation details to be aware of when testing:
- Data is stored in-memory and lost on restart.
- The POST handler currently sets the map key twice (`len(userCache)+1` and `user.ID`). Treat `id` as the source of truth when reading back.

## Quick cURL examples (PowerShell-friendly)

Create a user:

```pwsh
curl -X POST "http://localhost:8080/users" `
		 -H "Content-Type: application/json" `
		 -d '{"id":1,"name":"Alice","age":30}'
```

Fetch the user:

```pwsh
curl "http://localhost:8080/users/1"
```

Delete the user:

```pwsh
curl -X DELETE "http://localhost:8080/users/1"
```

## Using Postman

1) Start the server (see Run section).
2) In Postman, create an Environment:
	 - Variable: `baseUrl` = `http://localhost:8080`
3) Create a Collection (e.g., "Iris Web Server") with requests:
	 - Create User
		 - Method: POST
		 - URL: `{{baseUrl}}/users`
		 - Headers: `Content-Type: application/json`
		 - Body (raw/JSON):
			 ```json
			 { "id": 1, "name": "Alice", "age": 30 }
			 ```
		 - Tests (optional):
			 ```javascript
			 pm.test("Status is 204", function () {
				 pm.response.to.have.status(204);
			 });
			 ```
	 - Get User
		 - Method: GET
		 - URL: `{{baseUrl}}/users/1`
		 - Tests (optional):
			 ```javascript
			 pm.test("Status is 200", function () {
				 pm.response.to.have.status(200);
			 });
			 pm.test("Has name field", function () {
				 const data = pm.response.json();
				 pm.expect(data).to.have.any.keys('ID','name','Age');
			 });
			 ```
	 - Delete User
		 - Method: DELETE
		 - URL: `{{baseUrl}}/users/1`
		 - Tests (optional):
			 ```javascript
			 pm.test("Status is 204", function () {
				 pm.response.to.have.status(204);
			 });
			 ```

Tip: You can duplicate the requests and change `1` to other IDs for quick variations.

## Using VS Code REST clients (optional)

If you prefer working inside VS Code:

- REST Client extension: Create a file like `requests.http` with:
	```http
	### Create
	POST http://localhost:8080/users
	Content-Type: application/json

	{ "id": 1, "name": "Alice", "age": 30 }

	### Get
	GET http://localhost:8080/users/1

	### Delete
	DELETE http://localhost:8080/users/1
	```
- Thunder Client extension: Import the above 3 requests manually and set a `baseUrl` env var.

## Troubleshooting

- Port already in use: Change to a free port in `main.go` (e.g., `:8081`) and update your requests.
- 400 Bad Request on GET/DELETE: Ensure `id` is numeric in the path.
- 404 Not Found on GET: Make sure you created the user first with POST, and remember data resets on server restart.
- Running from the wrong directory: Prefer `go run ./Web_Server` from repo root.

## Next steps (optional improvements)

- Return 201 Created with a JSON body on POST.
- Normalize JSON field casing with struct tags.
- Add update (PUT/PATCH) support and basic validation tests.
