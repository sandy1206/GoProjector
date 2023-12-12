# GoProjector

Welcome to **GoProjector**, a comprehensive project management tool built with Golang. This efficient tool will help you manage tasks like never before. Secure your projects with JWT, store your data with MySQL, and deploy effortlessly with Docker.

## Description

GoProjector is a project management tool inspired by the efficiency and robustness of Golang. It uses Golang for the backend, JWT for secure authentication, and MySQL for reliable data storage. With Docker, you can easily deploy your project management solution to the cloud.

## Features

- **User Authentication:** Secure JWT-based authentication.
- **Project Management:** Create, update, delete, and manage projects.
- **Task Management:** Assign, update, delete, and track tasks within projects.
- **User Roles:** Admin and user roles with different permissions.
- **RESTful API:** Well-structured and documented REST API.
- **MySQL Database:** Reliable and scalable data storage.
- **Dockerized:** Easy deployment with Docker.

## Tech Stack

- **Golang:** Backend language.
- **JWT:** Authentication.
- **MySQL:** Database.
- **Docker:** Containerization for easy deployment.

## Installation

### Prerequisites

- Golang (version 1.16 or higher)
- MySQL
- Docker

### Steps

1. **Clone the repository:**
   ```sh
   git clone https://github.com/your-username/goprojector.git
   cd goprojector

2. **Set up the environment variables:**
   Create a `.env` file in the root directory and add the following variables:
   ```env
   DB_USER=<your-db-user>
   DB_PASSWORD=<your-db-password>
   DB_NAME=<your-db-name>
   DB_HOST=<your-db-host>
   JWT_SECRET=<your-jwt-secret>

3. **Run the application:**
   ```sh
   go run cmd/migrate/main.go

4. **Build the Docker image:**
   ```sh
   docker build -t goprojector .

5. **Run the Docker container:**
   ```sh
   docker-compose up --build

### API Endpoints
#### Authentication

#### Authentication
- `POST /api/auth/register`: Register a new user.
- `POST /api/auth/login`: Log in an existing user.
- `POST /api/auth/logout`: Log out the current user.

#### Project Management
- `GET /api/projects`: Get all projects.
- `GET /api/projects/{id}`: Get a specific project by ID.
- `POST /api/projects`: Create a new project.
- `PUT /api/projects/{id}`: Update an existing project.
- `DELETE /api/projects/{id}`: Delete a project.

#### Task Management
- `GET /api/projects/{projectId}/tasks`: Get all tasks within a project.
- `GET /api/projects/{projectId}/tasks/{taskId}`: Get a specific task by ID within a project.
- `POST /api/projects/{projectId}/tasks`: Create a new task within a project.
- `PUT /api/projects/{projectId}/tasks/{taskId}`: Update an existing task within a project.
- `DELETE /api/projects/{projectId}/tasks/{taskId}`: Delete a task within a project.

#### User Management
- `GET /api/users`: Get all users.
- `GET /api/users/{id}`: Get a specific user by ID.
- `PUT /api/users/{id}`: Update an existing user.
- `DELETE /api/users/{id}`: Delete a user.

## Contributing

We welcome contributions from the community to enhance GoProjector. To contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your changes to your forked repository.
5. Submit a pull request to the main repository.

Please ensure that your code follows the project's coding conventions and includes appropriate tests.

## License

GoProjector is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact

For any questions or inquiries, please contact me at arsudesandip4@gmail.com.
