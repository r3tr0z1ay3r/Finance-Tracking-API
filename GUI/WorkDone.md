# Work Done Log

This document details the implementation of the Finance Tracking GUI application.

## Project Structure & Implementation Details

### 1. `app.py`
- **Entry Point**: Initializes the `QApplication`.
- **Global Styling**:
    - Applied `Fusion` style for a consistent base.
    - Implemented a **Global Dark Theme** (Dracula-inspired palette) affecting `QMainWindow`, `QWidget`, and `QPushButton`.
    - Defined default font (`Segoe UI`) and base font size.

### 2. `ui_main.py`
- **MainWindow Class**: Core application logic and UI layout.
- **Layout**:
    - **Top Bar**: Month navigation (`<`, `>`) and dynamic Month/Year label. "Add Transaction" button.
    - **Central Table**: Displays transactions with columns: Date, Flow, Mode, Amount.
    - **Bottom Bar**: "Delete Transaction" button (enabled only when a row is selected).
- **Features**:
    - **Data Display**: Fetches and displays transactions for the selected month.
    - **Visual Cues**:
        - **Color Coding**: Income (Green), Expense (Red).
        - **Currency Formatting**: Amounts displayed in INR (₹).
    - **Interactivity**:
        - Single row selection.
        - Delete confirmation dialog with custom styling.
- **Styling Refinements**:
    - Custom `QTableWidget` styling (no grid, alternating row colors, custom header).
    - Removed focus outlines for a cleaner look.

### 3. `ui_dialogs.py`
- **AddTransactionDialog Class**: Modal dialog for inputting new transactions.
- **Form Fields**:
    - **Date**: `QDateEdit` with a popup calendar.
    - **Amount**: `QDoubleSpinBox` with INR prefix and empty initial state.
    - **Flow**: `QComboBox` (Income/Expense).
    - **Mode**: `QComboBox` (UPI, Cash, Transfer).
- **Styling**:
    - Comprehensive dark theme applied to all input widgets and the dialog itself.
    - Heavily customized `QCalendarWidget` to match the application's dark aesthetic.
    - Fixed spinbox arrow visibility.

### 4. `api_client.py`
- **API Integration**: Handles all communication with the backend (`localhost:8080`).
- **Methods**:
    - `get_transactions(year, month)`: GET `/trans/get/{month}/{year}`.
    - `add_transaction(...)`: POST `/trans/add`.
    - `delete_transaction(id)`: DELETE `/trans/del/{id}`.
- **Authentication**:
    - Implemented API Key authentication using `X-API-KEY` header.
    - **Configuration**: Loads `API_KEY` from `../config/.env` (relative to the script) using `python-dotenv`.
- **Data Handling**: Converts between API JSON format and UI-friendly dictionaries.

### 5. Configuration & Dependencies
- **`requirements.txt`**:
    - `PyQt5`: GUI framework.
    - `requests`: HTTP client for API calls.
    - `python-dotenv`: For loading environment variables.
- **`config/.env`**: Stores the API Key (located in sibling directory).
