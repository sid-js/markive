# Markive 🚀

Markive is a powerful and efficient Markdown Blog API built with Go-Chi router and GORM, enabling you to create, read, update, and delete blog posts seamlessly. It also integrates with PostgreSQL for robust data storage and Redis Cache for lightning-fast performance. Markive is designed to be fast, scalable, and easy to use, making it an excellent choice for developers looking to build markdown-based blogging platforms or applications.

## Features ✨

- **CRUD Operations**: Perform Create, Read, Update, and Delete operations on blog posts effortlessly.
- **PostgreSQL Integration**: Store your blog data securely in a PostgreSQL database.
- **Redis Cache**: Utilize Redis Cache to enhance the performance of your application.
- **Open Source**: Markive is open-source, allowing the community to contribute, enhance, and customize the project according to their needs.

## Getting Started 🚀

Follow these steps to get Markive up and running on your local machine:

### Prerequisites

- Go installed on your system
- PostgreSQL database set up
- Redis Cache installed and running

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/markive.git
   ```

3. Navigate to the project directory:
   ```sh
   cd markive
   ```

5. Install dependencies:
   ```sh
   go mod download
   ```

6. Configure your PostgreSQL and Redis connection in the `config/config.go` file.

7. Run the application:
   ```sh
   go run main.go
   ```

The API server will start running locally at `http://localhost:8080`.

## API Endpoints 🚪

- **Create Post**: `POST /api/posts`
- **Get All Posts**: `GET /api/posts`
- **Get Post by ID**: `GET /api/posts/{id}`
- **Update Post**: `PUT /api/posts/{id}`
- **Delete Post**: `DELETE /api/posts/{id}`

## Contribution Guidelines 🤝

We welcome contributions from the community! If you want to contribute to Markive, read the [Contribution Guidelines](CONTRIBUTING)



## License 📝

Markive is open-source software licensed under the [MIT License](LICENSE).

