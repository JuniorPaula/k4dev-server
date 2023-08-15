# K4Dev API - Knowledge for Developers

The **K4Dev API** is a backend project developed in Golang that serves as a platform for creating and managing articles, akin to a blog, while also providing user registration and management functionalities.

## Features

- **Article Management**: The API allows users to create, retrieve, update, and delete articles. Articles can be organized by categories and tags for easy navigation and discovery.
- **Categories**: Articles are categorized into different topics, making it easier for users to explore content that interests them.
- **User Registration**: Users can sign up for an account, providing their details such as username, email, and password. User authentication ensures secure access to the platform.

## Technologies Used

- Golang: The backend is developed using the Go programming language, known for its performance and efficiency.

- Database: The project uses a relational database MySQL to store article and user data and no-relational MongoDB for store stats.

## API Endpoints
#### Articles
- `GET /articles`: Retrieve a list of articles.
- `GET /articles?page=1`: Retrieve a list of articles using pagination.
- `GET /articles/{id}`: Retrieve details of a specific article.
- `GET /category/{categoryId}/articles`: Retrieve a articles by category id.
- `POST /articles`: Create a new article.
- `PUT /articles/{id}`: Update an existing article.
- `DELETE /articles/{id}`: Delete an article.

#### Category
- `GET /category`: Retrieve a list of category.
- `GET /articles?page=1`: Retrieve a list of articles using pagination.
- `GET /category/tree`: Retrieve details of category with tree.
- `GET /category/{categoryId}`: Retrieve details of a specific category.
- `POST /category`: Create a new category.
- `PUT /category/{id}`: Update an existing category.
- `DELETE /category/{id}`: Delete an category.

#### User
- `GET /users`: Retrieve a list of users.
- `GET /users/{id}`: Retrieve details of a specific users.
- `POST /users`: Create a new users.
- `POST /users/{id}/update-password`: Update user password.
- `POST /users/{id}/update-role`: Update user role.
- `PUT /users/{id}`: Update an existing user.
- `DELETE /users/{id}`: Delete an user.

#### Authentication
- `POST /signup`: Register a new user.
- `POST /login`: Authenticate and log in a user.
- `POST /validateToken`: Validate a specific user's token.

#### Stats
- `GET /stats`: Retrieve details of stats

## Setup

1. Clone this repository to your local machine.
2. Navigate to the project directory.
3. Set up your database and provide the connection details in the configuration file. Don't forget to run SQL script into `sql` folder in your favorite database client.
4. Run `go run ./cmd/main.go` to start the API server.

## Usage

- Access the API documentation and test the endpoints using tools like Postman or cURL.
- Register a new user using the `/signup` endpoint.
- Obtain an authentication token by logging in using the `/login` endpoint.
- Use the obtained token to interact with the article-related endpoints.

## Contribution

Contributions to enhance and improve the project are welcome! Feel free to submit issues or pull requests.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

## Contact

If you have any questions or suggestions, feel free to contact us at [luke.junnior@icloud.com](mailto:luke.junnior@icloud.com).

Thank you for using the K4Dev API to enhance knowledge sharing among developers!