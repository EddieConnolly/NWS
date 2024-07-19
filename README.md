# NWS

National Weather Service API Server

## Project Requirements

Write an HTTP server that serves the forecasted weather. Your server should expose an endpoint that:

    Accepts latitude and longitude coordinates
    Returns the short forecast for that area for Today (“Partly Cloudy” etc)
    Returns a characterization of whether the temperature is “hot”, “cold”, or “moderate” (use your discretion on mapping temperatures to each type)
    Use the National Weather Service API Web Service as a data source.

The purpose of this exercise is to provide a sample of your work that we can discuss together in the Technical Interview.

    We respect your time. Spend as long as you need, but we intend it to take around an hour.
    We do not expect a production-ready service, but you might want to comment on your shortcuts.
    The submitted project should build and have brief instructions so we can verify that it works.
    You may write in whatever language or stack you’re most comfortable in, but it’s recommended to use the language for the job you’re applying for.

## Dependencies

<https://github.com/gin-gonic/gin>

<https://www.weather.gov/documentation/services-web-api>

## Running

    go run cmd/main.go

## Testing

    go test ./...

## Usage

    GET localhost:8080/weather?lat=<>&long=<>

## Example

    get localhost:8080/weather?lat=34.052235&long=-118.243683

    {
        "times": [
            {
            "localTime": "2024-07-19T14:00:00-07:00",
            "shortForecast": "Sunny",
            "temperatureCharacterization": "Hot"
            },
            {
            "localTime": "2024-07-19T15:00:00-07:00",
            "shortForecast": "Sunny",
            "temperatureCharacterization": "Hot"
            },
            {
            "localTime": "2024-07-19T16:00:00-07:00",
            "shortForecast": "Sunny",
            "temperatureCharacterization": "Hot"
            },
            {
            "localTime": "2024-07-19T17:00:00-07:00",
            "shortForecast": "Sunny",
            "temperatureCharacterization": "Hot"
            },
            {
            "localTime": "2024-07-19T18:00:00-07:00",
            "shortForecast": "Mostly Clear",
            "temperatureCharacterization": "Warm"
            },
            {
            "localTime": "2024-07-19T19:00:00-07:00",
            "shortForecast": "Mostly Clear",
            "temperatureCharacterization": "Warm"
            },
            {
            "localTime": "2024-07-19T20:00:00-07:00",
            "shortForecast": "Mostly Clear",
            "temperatureCharacterization": "Warm"
            },
            {
            "localTime": "2024-07-19T21:00:00-07:00",
            "shortForecast": "Mostly Clear",
            "temperatureCharacterization": "Warm"
            }
        ]
    }
