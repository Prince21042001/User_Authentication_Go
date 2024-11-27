
# Project Overview

This project is a Go-based application designed for efficient API management and interaction. It uses Docker for containerization, MySQL for database management, and supports multiple endpoints for API functionalities.

## Features
- Modular architecture with controllers, models, and routes.
- MySQL integration with an `init.sql` script for initial setup.
- Dockerized environment for seamless deployment.
- RESTful API design.

---

## Installation Instructions

### Prerequisites
- Docker
- Docker Compose
- MySQL

### Steps
1. **Clone the Repository**
   ```bash
   git clone <repository-url>
   cd project
   ```

2. **Set Up the Environment**
   Ensure you have Docker and Docker Compose installed.

3. **Build and Run Containers**
   ```bash
   docker-compose up --build
   ```

4. **Verify the Setup**
   - Access the application at `http://localhost:<port>`.
   - Ensure the database is set up correctly by checking the MySQL container logs.

---

## API Documentation

### Base URL
`http://localhost:<port>`

### Endpoints

#### 1. **GET /endpoint-name**
- **Description**: Fetches data from the application.
- **Request**: None required.
- **Response**:
  ```json
  {
    "status": "success",
    "data": [...]
  }
  ```

#### 2. **POST /endpoint-name**
- **Description**: Adds new data to the application.
- **Request**:
  ```json
  {
    "field1": "value1",
    "field2": "value2"
  }
  ```
- **Response**:
  ```json
  {
    "status": "success",
    "message": "Data added successfully"
  }
  ```

#### 3. **PUT /endpoint-name/{id}**
- **Description**: Updates existing data.
- **Request**:
  ```json
  {
    "field1": "newValue1"
  }
  ```
- **Response**:
  ```json
  {
    "status": "success",
    "message": "Data updated successfully"
  }
  ```

#### 4. **DELETE /endpoint-name/{id}**
- **Description**: Deletes data by ID.
- **Request**: None required.
- **Response**:
  ```json
  {
    "status": "success",
    "message": "Data deleted successfully"
  }
  ```

---

## Database
- **MySQL Initialization**: The `init.sql` file contains the schema and initial data setup for the application.
- Ensure the `docker-compose.yml` file links the MySQL container properly.

---

## Technologies Used
- **Go**: Backend language.
- **Docker**: For containerization.
- **MySQL**: Database management.
- **HTML Templates**: Frontend templates.

## Contribution
Feel free to fork this project, raise issues, or submit pull requests.

---

## License
This project is licensed under the MIT License.

## Team Members

- Prince Patel
- Nidhi Patel
- Ethili Sundarvel
- Sheetal Shivkumarn
