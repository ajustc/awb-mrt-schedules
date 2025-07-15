# AWB MRT Schedules

A Go service that provides Jakarta MRT station schedules and information.

## Features

- Get all MRT stations information
- Get schedule by station ID
- Real-time schedule updates

## API Endpoints

### Get All Stations
```
GET /stations
```
Returns a list of all MRT stations with their IDs and names.

### Get Station Schedule
```
GET /stations/{id}
```
Returns the upcoming schedule for a specific station, showing departure times for both Lebak Bulus and Bundaran HI directions.

## Setup

1. Clone the repository
2. Install dependencies:
```bash
go mod download
```
3. Run the service:
```bash
go run main.go
```

## Project Structure

```
├── common/          # Common utilities
│   ├── client/      # HTTP client implementation
│   └── response/    # Response handling
└── modules/
    └── station/     # Station-related functionality
        ├── dto.go   # Data transfer objects
        ├── router.go # Route definitions
        └── service.go # Business logic
```