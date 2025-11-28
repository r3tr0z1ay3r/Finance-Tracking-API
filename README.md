# Finance Tracking API & GUI

A comprehensive finance tracking solution featuring a robust Go backend API and a modern, dark-themed Python PyQt5 GUI.

## Features

*   **Transaction Management**: Add, view, and delete transactions seamlessly.
*   **Monthly Filtering**: Navigate through transaction history by month and year.
*   **Visual Insights**:
    *   **Flow Tracking**: Distinct visual cues for Income (Green) and Expenses (Red).
    *   **Mode Tracking**: Record transaction modes (UPI, Cash, Transfer).
*   **Modern UI**: A polished, dark-themed interface built with PyQt5 for a premium user experience.
*   **Robust Backend**: High-performance REST API built with Go and SQLite.

## Technology Stack

### Backend
*   **Language**: Go (Golang) 1.25+
*   **Router**: Gorilla Mux
*   **Database**: SQLite (via `modernc.org/sqlite`)
*   **Architecture**: RESTful API

### Frontend
*   **Language**: Python 3.x
*   **Framework**: PyQt5
*   **Styling**: Custom Dark Theme (Dracula-inspired)

## Prerequisites

*   [Go](https://go.dev/dl/) (version 1.25 or higher)
*   [Python](https://www.python.org/downloads/) (version 3.8 or higher)

## Installation & Setup

### 1. Backend Setup

1.  Navigate to the project root directory.
2.  Install Go dependencies:
    ```bash
    go mod tidy
    ```
3.  Run the server:
    ```bash
    go run cmd/Server/main.go
    ```
    The server will start listening on `http://localhost:8080`.

### 2. Frontend Setup

1.  Navigate to the `GUI` directory:
    ```bash
    cd GUI
    ```
2.  Create a virtual environment (optional but recommended):
    ```bash
    python -m venv .venv
    # Windows
    .venv\Scripts\activate
    # Linux/Mac
    source .venv/bin/activate
    ```
3.  Install Python dependencies:
    ```bash
    pip install -r requirements.txt
    ```
4.  **Configuration**:
    Ensure you have a `.env` file in the `Config` directory (or configured as required by `api_client.py`) containing your API key if authentication is enabled.
    *   *Note: The current `api_client.py` looks for `../config/.env` relative to the GUI folder.*

5.  Run the application:
    ```bash
    python app.py
    ```

## API Endpoints

The backend exposes the following REST endpoints:

*   `GET /trans/get/{month}/{year}`: Retrieve transactions for a specific month and year.
*   `POST /trans/add`: Add a new transaction.
    *   Payload: JSON object with transaction details.
*   `DELETE /trans/del/{id}`: Delete a transaction by ID.

## Project Structure

*   `cmd/Server/`: Contains the main entry point for the Go backend.
*   `GUI/`: Contains the Python PyQt5 application source code.
*   `Internal/`: Internal Go packages for API logic and Database handling.
*   `Config/`: Configuration files (e.g., `.env`).

## Future Enhancements

*   **User-Based System**: Implement multi-user support to track transactions for different users independently.
*   **Spending Analytics**: Develop analytics to track and visualize user spending habits over time.
*   **Financial Insights**: Provide analytical insights to improve financial awareness and budgeting.
*   **Web-Based GUI**: Develop a web interface for multi-device access and host the API online.
