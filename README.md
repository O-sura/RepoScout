# RepoScout

[![Tool Set](https://skillicons.dev/icons?i=golang,react,java,docker,jenkins,github)](https://skillicons.dev)

A generous attempt which helps developers to connect with their desired open source project :D

# Project Structure

```plaintext
reposcout/
├── backend/
│   ├── discovery-service/        # Golang service for repository discovery
│   │   ├── main.go               # Entry point for the service
│   │   ├── go.mod                # Go module file
│   │   ├── go.sum                # Go dependency lock file
│   │   ├── internal/             # Internal packages for business logic
│   │   ├── api/                  # API routes and handlers
│   │   └── utils/                # Utility functions
│   ├── preferences-service/      # Java service for managing user preferences
│   │   ├── src/
│   │   │   ├── main/
│   │   │   │   ├── java/
│   │   │   │   │   ├── com/
│   │   │   │   │   │   ├── example/
│   │   │   │   │   │   │   ├── preferences/   # Java packages
│   │   │   │   │   │   │   │   ├── controllers/ # Controllers for endpoints
│   │   │   │   │   │   │   │   ├── services/    # Business logic
│   │   │   │   │   │   │   │   ├── models/      # Data models
│   │   │   │   │   │   │   │   └── repositories/ # Database access
│   │   │   └── resources/       # Configuration files (e.g., `application.yml`)
│   │   └── pom.xml              # Maven configuration file
├── frontend/
│   ├── public/                  # Static files (e.g., index.html, favicon)
│   ├── src/                     # React application source code
│   │   ├── components/          # Reusable components
│   │   ├── pages/               # Pages (e.g., home, repository details)
│   │   ├── services/            # API calls to backend
│   │   ├── App.js               # Main React component
│   │   ├── index.js             # Entry point for React
│   │   └── styles/              # CSS or SCSS files
│   ├── package.json             # Node.js dependencies
│   └── vite.config.js           # Vite or Webpack config (if using Vite/React)
├── docker/
│   ├── discovery-service/       # Dockerfile for Golang service
│   ├── preferences-service/     # Dockerfile for Java service
│   └── frontend/                # Dockerfile for React frontend
├── docker-compose.yml           # Docker Compose configuration
├── README.md                    # Project documentation
└── .gitignore                   # Ignore unnecessary files
```

## Setup/Installation

## Usage

#### Repository Discovery Service

   - **Endpoints**:
     - `GET /repositories`: Fetch repositories matching query parameters (e.g., language, tags).
     - `GET /repository/{id}`: Fetch details of a specific repository.


#### User Preferences Service
   - **Endpoints**:
     - `POST /user`: Create or update user preferences.
     - `GET /user/{id}`: Retrieve user preferences.
     - `GET /user/{id}/matches`: Get recommended repositories for the user.

## How to Contribute

