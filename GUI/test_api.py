from api_client import APIClient
import datetime

def test_api():
    client = APIClient()
    
    print("Testing Add Transaction...")
    # Add a test transaction
    success = client.add_transaction(
        date_str=datetime.date.today().strftime("%Y-%m-%d"),
        flow="Expense",
        mode="TestMode",
        amount=123.45
    )
    print(f"Add Transaction Success: {success}")
    
    print("\nTesting Get Transactions...")
    # Get transactions for current month
    today = datetime.date.today()
    transactions = client.get_transactions(today.year, today.month)
    print(f"Found {len(transactions)} transactions.")
    
    found_t = None
    for t in transactions:
        print(f" - {t['date']} {t['flow']} {t['mode']} {t['amount']} (ID: {t['id']})")
        if t['mode'] == "TestMode" and t['amount'] == 123.45:
            found_t = t
            
    if found_t:
        print(f"\nFound our test transaction with ID: {found_t['id']}")
        print("Testing Delete Transaction...")
        del_success = client.delete_transaction(found_t['id'])
        print(f"Delete Transaction Success: {del_success}")
        
        # Verify deletion
        transactions_after = client.get_transactions(today.year, today.month)
        still_there = any(t['id'] == found_t['id'] for t in transactions_after)
        print(f"Transaction still exists after delete: {still_there}")
    else:
        print("\nCould not find the test transaction to delete.")

if __name__ == "__main__":
    test_api()
