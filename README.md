# URL Shortener Service

This is a distributed, cloud-native URL Shortener Service built with Go.

## API Endpoints

### Shorten URL

- **URL**: `/shorten`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "url": "https://www.example.com"
  }
  ```
