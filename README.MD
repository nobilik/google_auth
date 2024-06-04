# Go Frontend and Backend Project

This project represents separate frontend and backend written in Go and packaged into Docker containers. The project utilizes standard Go packages and provides login authentication via username and password, as well as Google OAuth.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/nobilik/google_auth.git
    ```

2. Add a `.env` file to the root of the project with the following parameters:

    ```
    GOOGLE_CLIENT_ID=your_google_id
    GOOGLE_CLIENT_SECRET=your_google_secret
    API_HOST=http://backend:3000
    HOST=http://localhost:8080
    MYSQL_ROOT_PASSWORD=your_root_password
    MYSQL_DATABASE=your_database
    MYSQL_USER=your_user
    MYSQL_PASSWORD=your_password
    ```

3. Build and run the project using Docker Compose:

    ```bash
    docker-compose up --build
    ```

## Authors

- Author: nobilik

## License

This project is licensed under the MIT License.
