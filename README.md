# Structure of the Service

**config**
Initializes and validates environment variables from the .env file, ensuring essential configurations are set for different environments.

**docs**
Contains API documentation (Swagger) for easy access to endpoint details, request/response formats, and usage examples, aiding developers in integration and testing.

**dto**
Defines Data Transfer Objects (DTOs) for request and response structures, maintaining consistency in data handling across the API.

**e2e**
Includes integration and load testing code using k6, validating API performance and reliability before deployment.

**entities**
Houses Go structs that represent database tables, facilitating structured data manipulation and strong typing in the application.

**handlers**
Contains controllers for processing HTTP requests, such as authentication (auth.go) and profile management (profile.go), which route data to the appropriate services.

**migrations**
Holds SQL scripts for creating and modifying database schemas, allowing version control for database structure changes.

**pkg**
Includes utility functions for common tasks like database connections, error handling, and logging, promoting code reuse and modularity.

**services**
Encapsulates core business logic, processing data and enforcing business rules before interacting with handlers and repositories.

# How to Run the Service

1. **Set Up the PostgreSQL Database Container**  
   Begin by creating and starting a PostgreSQL database container using Docker. Execute the following command in your terminal:

   ```bash
   make up-db
   ```

   This command initializes a PostgreSQL instance within a Docker container, preparing the database for your application.

2. **Copy Environment Configuration**  
   Copy the template environment file to create your environment configuration:

   ```bash
   cp .env.template .env
   ```

   This step is crucial as it sets up the necessary environment variables that the application requires to run correctly, such as database connection details.

3. **Apply Database Migrations**  
   Next, run the migrations to set up the required tables in the PostgreSQL database. Use the following command:

   ```bash
   make migrate-up
   ```

   This command applies any existing migrations, ensuring your database schema is up to date.

4. **Start the Backend Service**  
   To launch the backend service, execute:

   ```bash
   make run
   ```

   This command starts your backend application, which will connect to the PostgreSQL database and be ready to handle incoming requests.

5. **Run Integration Tests**  
   After the backend service is running, you can test its functionality and performance. Ensure you have `k6` installed. If you haven’t installed it yet, follow the instructions available at [Install k6](https://grafana.com/docs/k6/latest/set-up/install-k6/). Once `k6` is set up, run the integration tests using:

   ```bash
   make test
   ```

   This command executes integration tests with `k6`, allowing you to verify that the service behaves as expected under load.

By following these steps, you will successfully set up and run the backend service, ensuring your database is correctly configured and tested.

### [PDF Doc](https://drive.google.com/file/d/1jfJJotc-ilG3SyCXAZgF2QBBsd_Sov3o/view)