from PyQt5.QtWidgets import (QDialog, QVBoxLayout, QFormLayout, QLineEdit, 
                             QDateEdit, QDoubleSpinBox, QDialogButtonBox, QComboBox,QLabel)
from PyQt5.QtCore import QDate

class AddTransactionDialog(QDialog):
    def __init__(self, parent=None):
        super().__init__(parent)
        self.setWindowTitle("Add Transaction")
        self.setModal(True)
        self.resize(400, 250)
        self.setStyleSheet("""
            QDialog {
                background-color: #282a36;
                color: #f8f8f2;
            }
            QLabel {
                color: #f8f8f2;
                font-size: 14px;
            }
            QLineEdit, QComboBox, QDateEdit, QDoubleSpinBox {
                background-color: #44475a;
                color: #f8f8f2;
                border: 1px solid #6272a4;
                border-radius: 5px;
                padding: 5px;
                font-size: 14px;
            }
            QComboBox QAbstractItemView {
                background-color: #44475a;
                color: #f8f8f2;
                selection-background-color: #6272a4;
                selection-color: #f8f8f2;
                border: 1px solid #6272a4;
            }
            /* Fix SpinBox Arrows */
            QDoubleSpinBox::up-button, QDoubleSpinBox::down-button {
                width: 20px;
                background-color: #6272a4;
                border: none;
                border-radius: 2px;
            }
            QDoubleSpinBox::up-button:hover, QDoubleSpinBox::down-button:hover {
                background-color: #50fa7b;
            }
            QDoubleSpinBox::up-arrow {
                image: none;
                width: 0; 
                height: 0; 
                border-left: 4px solid transparent;
                border-right: 4px solid transparent;
                border-bottom: 4px solid #f8f8f2;
            }
            QDoubleSpinBox::down-arrow {
                image: none;
                width: 0; 
                height: 0; 
                border-left: 4px solid transparent;
                border-right: 4px solid transparent;
                border-top: 4px solid #f8f8f2;
            }
            
            /* Calendar Widget Styling */
            QCalendarWidget QToolButton {
                color: #f8f8f2;
                background-color: #44475a;
                icon-size: 20px;
            }
            QCalendarWidget QMenu {
                background-color: #44475a;
                color: #f8f8f2;
            }
            QCalendarWidget QSpinBox {
                background-color: #44475a;
                color: #f8f8f2;
                selection-background-color: #6272a4;
            }
            QCalendarWidget QWidget#qt_calendar_navigationbar {
                background-color: #44475a;
            }
            QCalendarWidget QAbstractItemView:enabled {
                color: #f8f8f2;
                background-color: #282a36;
                selection-background-color: #6272a4;
                selection-color: #f8f8f2;
                outline: 0;
            }
            QCalendarWidget QAbstractItemView:disabled {
                color: #6272a4;
            }
            /* Fix Calendar Header (Day Names) */
            QCalendarWidget QWidget {
                alternate-background-color: #44475a; 
            }
            QCalendarWidget QTableView {
                background-color: #282a36;
                alternate-background-color: #44475a;
            }
            
            QPushButton {
                background-color: #6272a4;
                color: #f8f8f2;
                border: none;
                border-radius: 5px;
                padding: 8px 16px;
                font-size: 14px;
            }
            QPushButton:hover {
                background-color: #50fa7b;
                color: #282a36;
            }
        """)

        layout = QVBoxLayout(self)
        
        # Add description label
        description_label = QLabel("Enter the details of the new transaction below:")
        description_label.setStyleSheet("color: #bd93f9; font-style: italic; margin-bottom: 10px;")
        layout.addWidget(description_label)

        form_layout = QFormLayout()
        form_layout.setSpacing(15)

        self.date_edit = QDateEdit(QDate.currentDate())
        self.date_edit.setCalendarPopup(True)
        self.date_edit.setDisplayFormat("yyyy-MM-dd")
        
        self.amount_spin = QDoubleSpinBox()
        self.amount_spin.setRange(0, 1000000)
        self.amount_spin.setDecimals(2)
        self.amount_spin.setPrefix("₹")
        self.amount_spin.setSpecialValueText(" ") # Make it appear empty when 0

        self.flow_combo = QComboBox()
        self.flow_combo.addItems(["Income", "Expense"])

        self.mode_combo = QComboBox()
        self.mode_combo.addItems(["UPI", "Cash", "Transfer"])

        form_layout.addRow("Date:", self.date_edit)
        form_layout.addRow("Amount:", self.amount_spin)
        form_layout.addRow("Flow:", self.flow_combo)
        form_layout.addRow("Mode:", self.mode_combo)

        layout.addLayout(form_layout)

        self.buttons = QDialogButtonBox(QDialogButtonBox.Ok | QDialogButtonBox.Cancel)
        self.buttons.accepted.connect(self.accept)
        self.buttons.rejected.connect(self.reject)
        layout.addWidget(self.buttons)

    def get_data(self):
        return {
            "date": self.date_edit.date().toString("yyyy-MM-dd"),
            "amount": self.amount_spin.value(),
            "flow": self.flow_combo.currentText(),
            "mode": self.mode_combo.currentText()
        }
