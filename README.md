# Hotel Reservation Backend API

This is a backend API for a hotel reservation system built in Go. The API allows users to manage hotels, rooms, and bookings. It also includes user authentication and authorization.

## Packages Used

- **github.com/gofiber/fiber/v2**: Web framework for building REST APIs.
- **github.com/golang-jwt/jwt**: User authentication and authorization using JWT.
- **github.com/joho/godotenv**: Loads environment variables from a `.env` file.
- **go.mongodb.org/mongo-driver/mongo**: MongoDB driver for Go.

## Architecture Diagram

<img alt="Architecture Diagram" src="/assets/architecture.png">

## Endpoint Flow

<img alt="Endpoint Flow Diagram" src="/assets/endpoint.png">

## Endpoints

### Authentication

- **POST /api/auth**: Authenticates a user and returns a JWT token.
  - **Handler**: `HandleAuthentication`
  - **Implementation**: Validates user credentials and generates a JWT token using `CreateTokenFromUser`

### User Management

- **POST /api/v1/user**: Creates a new user.
  - **Handler**: `HandlePostUser`
  - **Implementation**: Parses user data, validates it, and stores it in the database.

- **GET /api/v1/users**: Retrieves all users.
  - **Handler**: `HandleGetUsers`
  - **Implementation**: Fetches all users from the database.

- **GET /api/v1/user/:id**: Retrieves a user by ID.
  - **Handler**: `HandleGetUser`
  - **Implementation**: Fetches a user by ID from the database.

- **PUT /api/v1/user/:id**: Updates a user by ID.
  - **Handler**: `HandleUpdateUser`
  - **Implementation**: Updates user data in the database.

- **DELETE /api/v1/user/:id**: Deletes a user by ID.
  - **Handler**: `HandleDeleteUser`
  - **Implementation**: Deletes a user from the database.

### Hotel Management

- **GET /api/v1/hotels**: Retrieves all hotels.
  - **Handler**: `GetHotels`
  - **Implementation**: Fetches all hotels from the database.

- **GET /api/v1/hotel/:id**: Retrieves a hotel by ID.
  - **Handler**: `GetHotel`
  - **Implementation**: Fetches a hotel by ID from the database.

- **GET /api/v1/hotel/:id/rooms**: Retrieves all rooms for a hotel.
  - **Handler**: `GetRooms`
  - **Implementation**: Fetches all rooms for a specific hotel from the database.

### Room Management

- **POST /api/v1/room/:id/book**: Books a room.
  - **Handler**: `HandleBookRoom`
  - **Implementation**: Validates booking data, checks room availability, and stores the booking in the database.

### Booking Management

- **GET /api/v1/booking/:id**: Retrieves a booking by ID.
  - **Handler**: `HandleGetBooking`
  - **Implementation**: Fetches a booking by ID from the database.

- **GET /api/v1/bookings**: Retrieves all bookings.
  - **Handler**: `HandleGetBookings`
  - **Implementation**: Fetches all bookings from the database.

- **GET /api/v1/booking/:id/cancel**: Cancels a booking by ID.
  - **Handler**: `HandleCancelBooking`
  - **Implementation**: Cancels a booking by ID in the database.

### Admin Endpoints

- **GET /api/v1/admin/bookings**: Retrieves all bookings (admin only).
  - **Handler**: `HandleGetBookings`
  - **Implementation**: Fetches all bookings from the database.

## Getting Started

1. **Clone the repository**:
    ```sh
    git clone https://github.com/yourusername/hotel-reservation.git
    cd hotel-reservation
    ```

2. **Install dependencies**:
    ```sh
    go mod tidy
    ```

3. **Set up environment variables**:
    Create a `.env` file in the root directory and add the following variables:
    ```
    JWT_SECRET=your_jwt_secret
    DB_URI=your_mongodb_uri
    DB_NAME=your_database_name
    ```

4. **Use Make**:
 - **Build the application**:
    ```sh
    make build
    ```

- **Run the application**:
    ```sh
    make run
    ```

- **Seed the database**:
    ```sh
    make seed
    ```

- **Run tests**:
    ```sh
    make test
    ```