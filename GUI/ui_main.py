from PyQt5.QtWidgets import (QMainWindow, QWidget, QVBoxLayout, QHBoxLayout, 
                             QPushButton, QLabel, QTableWidget, QTableWidgetItem, 
                             QHeaderView, QMessageBox, QAbstractItemView, QDialog)
from PyQt5.QtCore import Qt, QDate
from api_client import APIClient
from ui_dialogs import AddTransactionDialog

class MainWindow(QMainWindow):
    def __init__(self):
        super().__init__()
        self.setWindowTitle("Finance Tracker")
        self.resize(800, 600)

        self.api_client = APIClient()
        self.current_date = QDate.currentDate()

        self.setup_ui()
        self.load_transactions()

    def setup_ui(self):
        central_widget = QWidget()
        self.setCentralWidget(central_widget)
        main_layout = QVBoxLayout(central_widget)
        main_layout.setSpacing(20)
        main_layout.setContentsMargins(20, 20, 20, 20)

        # Top Bar
        top_bar = QHBoxLayout()
        
        self.prev_month_btn = QPushButton("<")
        self.prev_month_btn.setFixedSize(40, 40)
        self.prev_month_btn.clicked.connect(self.prev_month)
        
        self.month_label = QLabel()
        self.month_label.setAlignment(Qt.AlignCenter)
        self.month_label.setStyleSheet("font-size: 18px; font-weight: bold; color: #f8f8f2;")
        self.month_label.setMinimumWidth(200)
        self.update_month_label()

        self.next_month_btn = QPushButton(">")
        self.next_month_btn.setFixedSize(40, 40)
        self.next_month_btn.clicked.connect(self.next_month)

        self.add_btn = QPushButton("Add Transaction")
        self.add_btn.setCursor(Qt.PointingHandCursor)
        self.add_btn.setStyleSheet("""
            QPushButton {
                background-color: #50fa7b;
                color: #282a36;
                font-weight: bold;
                padding: 10px 20px;
                border-radius: 5px;
            }
            QPushButton:hover {
                background-color: #40e06b;
            }
        """)
        self.add_btn.clicked.connect(self.add_transaction)

        top_bar.addWidget(self.prev_month_btn)
        top_bar.addWidget(self.month_label)
        top_bar.addWidget(self.next_month_btn)
        top_bar.addStretch()
        top_bar.addWidget(self.add_btn)

        main_layout.addLayout(top_bar)

        # Transaction Table
        self.table = QTableWidget()
        self.table.setColumnCount(5)
        self.table.setHorizontalHeaderLabels(["Date", "Flow", "Mode", "Amount", "ID"])
        self.table.horizontalHeader().setSectionResizeMode(QHeaderView.Stretch)
        self.table.setSelectionBehavior(QAbstractItemView.SelectRows)
        self.table.setSelectionMode(QAbstractItemView.SingleSelection)
        self.table.setEditTriggers(QAbstractItemView.NoEditTriggers)
        self.table.hideColumn(4) # Hide ID column
        self.table.verticalHeader().setVisible(False)
        self.table.setShowGrid(False)
        self.table.setAlternatingRowColors(True)
        self.table.setStyleSheet("""
            QTableWidget {
                background-color: #282a36;
                color: #f8f8f2;
                border: none;
                gridline-color: #44475a;
                selection-background-color: #f8f8f2;
                selection-color: #282a36;
                selection-color: #282a36;
                alternate-background-color: #44475a;
                outline: 0;
            }
            QHeaderView::section {
                background-color: #44475a;
                color: #f8f8f2;
                padding: 10px;
                border: none;
                font-weight: bold;
            }
            QTableWidget::item {
                padding: 10px;
                border-bottom: 1px solid #44475a;
            }
            QTableWidget::item:selected {
                background-color: #f8f8f2;
                color: #282a36;
            }
        """)
        self.table.itemSelectionChanged.connect(self.update_buttons)
        
        main_layout.addWidget(self.table)

        # Bottom Bar
        bottom_bar = QHBoxLayout()
        self.delete_btn = QPushButton("Delete Transaction")
        self.delete_btn.setEnabled(False)
        self.delete_btn.setCursor(Qt.PointingHandCursor)
        self.delete_btn.setStyleSheet("""
            QPushButton {
                background-color: #ff5555;
                color: #f8f8f2;
                font-weight: bold;
                padding: 10px 20px;
                border-radius: 5px;
            }
            QPushButton:hover {
                background-color: #ff4444;
            }
            QPushButton:disabled {
                background-color: #44475a;
                color: #6272a4;
            }
        """)
        self.delete_btn.clicked.connect(self.delete_transaction)
        
        bottom_bar.addStretch()
        bottom_bar.addWidget(self.delete_btn)

        main_layout.addLayout(bottom_bar)

    def update_month_label(self):
        self.month_label.setText(self.current_date.toString("MMMM yyyy"))

    def prev_month(self):
        self.current_date = self.current_date.addMonths(-1)
        self.update_month_label()
        self.load_transactions()

    def next_month(self):
        self.current_date = self.current_date.addMonths(1)
        self.update_month_label()
        self.load_transactions()

    def load_transactions(self):
        self.table.setRowCount(0)
        self.table.clearSelection() # Fix stuck highlight
        transactions = self.api_client.get_transactions(self.current_date.year(), self.current_date.month())
        
        self.table.setRowCount(len(transactions))
        for i, t in enumerate(transactions):
            self.table.setItem(i, 0, QTableWidgetItem(t['date']))
            self.table.setItem(i, 1, QTableWidgetItem(t['flow']))
            self.table.setItem(i, 2, QTableWidgetItem(t['mode']))
            
            amount_item = QTableWidgetItem(f"₹{t['amount']:.2f}")
            amount_item.setTextAlignment(Qt.AlignRight | Qt.AlignVCenter)
            
            # Color code amount based on flow
            if t['flow'] == 'Income':
                amount_item.setForeground(Qt.green)
            else:
                amount_item.setForeground(Qt.red)
                
            self.table.setItem(i, 3, amount_item)
            
            self.table.setItem(i, 4, QTableWidgetItem(t['id']))

    def update_buttons(self):
        selected = self.table.selectedItems()
        self.delete_btn.setEnabled(len(selected) > 0)

    def add_transaction(self):
        dialog = AddTransactionDialog(self)
        if dialog.exec_() == QDialog.Accepted:
            data = dialog.get_data()
            self.api_client.add_transaction(data['date'], data['flow'], data['mode'], data['amount'])
            self.load_transactions()

    def delete_transaction(self):
        row = self.table.currentRow()
        if row < 0:
            return
            
        t_id = self.table.item(row, 4).text()
        

        msg = QMessageBox(self)
        msg.setWindowTitle("Delete Transaction")
        msg.setText("Are you sure you want to delete this transaction?")
        msg.setInformativeText("This action cannot be undone.")
        msg.setIcon(QMessageBox.Warning)
        msg.setStandardButtons(QMessageBox.Yes | QMessageBox.No)
        msg.setDefaultButton(QMessageBox.No)
        msg.setStyleSheet("""
            QMessageBox {
                background-color: #282a36;
                color: #f8f8f2;
            }
            QLabel {
                color: #f8f8f2;
            }
            QPushButton {
                background-color: #44475a;
                color: #f8f8f2;
                border: none;
                border-radius: 5px;
                padding: 5px 15px;
            }
            QPushButton:hover {
                background-color: #6272a4;
            }
        """)
        
        if msg.exec_() == QMessageBox.Yes:
            self.api_client.delete_transaction(t_id)
            self.load_transactions()
