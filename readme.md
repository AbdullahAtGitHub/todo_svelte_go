# Todo Svelte Go Project

## Getting Started

Follow these steps to set up and run the project:

### Prerequisites

- Docker
- Go
- Node.js and npm

### Setup

1. **Start the Database**

   Start the database using Docker Compose which is present in todo_backend folder:

   ```sh
   docker-compose up -d
   ```

2. **Initialize the Database**

   Initialize the database tables using the provided SQL script:

   ```sh
   psql -h localhost -U <username> -d <database> -f init_todo.sql
   ```

   Replace `<username>` and `<database>` with the credentials specified in the Docker file.

3. **Run the Backend**

   Navigate to the backend folder and run the Go application:

   ```sh
   cd backend
   go run main.go
   ```

4. **Run the Frontend**

   Navigate to the frontend folder and start the Svelte application:

   ```sh
   cd frontend
   npm install
   npm run dev
   ```

### Project Structure

- `backend/` - Contains the Go backend code.
- `frontend/` - Contains the Svelte frontend code.
- `init_todo.sql` - SQL script to initialize the database.

### Additional Information

- Ensure Docker is running before starting the database.
- The database credentials are specified in the Docker file.

Enjoy building with Todo Svelte Go!
