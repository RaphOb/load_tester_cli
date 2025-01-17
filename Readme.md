
{
  "routes": [
    {
      "method": "GET",
      "url": "http://127.0.0.1:8001/test/sync_async",
      "headers": {
        "Authorization": "Bearer <token>"
      },
      "body": null,
      "description": "Retrieve a list of users"
    },
    {
      "method": "POST",
      "url": "/users",
      "headers": {
        "Content-Type": "application/json",
        "Authorization": "Bearer <token>"
      },
      "body": {
        "name": "John Doe",
        "email": "john.doe@example.com"
      },
      "description": "Create a new user"
    },
    {
      "method": "PUT",
      "url": "/users/{id}",
      "headers": {
        "Content-Type": "application/json",
        "Authorization": "Bearer <token>"
      },
      "body": {
        "name": "Updated Name"
      },
      "description": "Update user information"
    }
  ]
}