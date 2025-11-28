import requests
import datetime
import json
import os
from dotenv import load_dotenv

# Load .env from ../config/.env
dotenv_path = os.path.join(os.path.dirname(os.path.dirname(os.path.abspath(__file__))), 'config', '.env')
load_dotenv(dotenv_path)

class APIClient:
    def __init__(self, base_url="http://localhost:8080"):
        self.base_url = base_url
        self.headers = {
            "X-API-KEY": os.getenv("API_KEY")
        }

    def get_transactions(self, year, month):
        """
        Fetch transactions for a specific year and month.
        GET /trans/get/{month}/{year}
        """
        try:
            url = f"{self.base_url}/trans/get/{month}/{year}"
            response = requests.get(url, headers=self.headers)
            response.raise_for_status()
            
            data = response.json()
            if not data:
                return []
                
            # Transform API data to UI format
            # API returns: {ID, Time, Amt, Flow, Mode}
            # UI expects: {id, date, flow, mode, amount}
            transactions = []
            for item in data:
                # Parse time string to date string YYYY-MM-DD
                #  e.g. "2023-10-27T10:00:00Z"
                try:
                    dt = datetime.datetime.fromisoformat(item['time'].replace('Z', '+00:00'))
                    date_str = dt.strftime("%Y-%m-%d")
                except ValueError:
                    # Fallback or handle error
                    date_str = item['time'][:10] 

                transactions.append({
                    "id": str(item['id']), # Convert int ID to string for UI consistency
                    "date": date_str,
                    "flow": item['flow'],
                    "mode": item['mode'],
                    "amount": item['amt']
                })
            
            # Sort by date descending
            transactions.sort(key=lambda x: x['date'], reverse=True)
            return transactions
            
        except requests.exceptions.RequestException as e:
            print(f"Error fetching transactions: {e}")
            return []
        except Exception as e:
            print(f"Error processing transactions: {e}")
            return []

    def add_transaction(self, date_str, flow, mode, amount):
        """
        Add a new transaction.
        POST /trans/add
        Body: {ID, Time, Amt, Flow, Mode}
        """
        try:
            url = f"{self.base_url}/trans/add"
            
            # Convert date_str (YYYY-MM-DD) to ISO8601 time string
            # We'll set time to current time or 00:00:00
            dt = datetime.datetime.strptime(date_str, "%Y-%m-%d")
            # Add current time component if it's today, otherwise 12:00
            now = datetime.datetime.now()
            if dt.date() == now.date():
                dt = dt.replace(hour=now.hour, minute=now.minute, second=now.second)
            else:
                dt = dt.replace(hour=12, minute=0, second=0)
                
            time_str = dt.isoformat() + "Z" # Simple Z suffix for UTC/Zulu if backend expects it, or just isoformat
            
            payload = {
                "id": 0, # Backend should assign ID
                "time": time_str,
                "amt": float(amount),
                "flow": flow,
                "mode": mode
            }
            
            response = requests.post(url, json=payload, headers=self.headers)
            response.raise_for_status()
            return True
            
        except requests.exceptions.RequestException as e:
            print(f"Error adding transaction: {e}")
            return False

    def delete_transaction(self, transaction_id):
        """
        Delete a transaction by ID.
        DELETE /trans/del/{id}
        """
        try:
            url = f"{self.base_url}/trans/del/{transaction_id}"
            response = requests.delete(url, headers=self.headers)
            response.raise_for_status()
            return True
        except requests.exceptions.RequestException as e:
            print(f"Error deleting transaction: {e}")
            return False
