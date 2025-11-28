import sys
from PyQt5.QtWidgets import QApplication
from ui_main import MainWindow

def main():
    app = QApplication(sys.argv)
    
    # Set application style
    app.setStyle("Fusion")
    
    # Global Dark Theme
    app.setStyleSheet("""
        QMainWindow {
            background-color: #282a36;
        }
        QWidget {
            color: #f8f8f2;
            font-family: 'Segoe UI', sans-serif;
            font-size: 14px;
        }
        QPushButton {
            background-color: #44475a;
            color: #f8f8f2;
            border: none;
            border-radius: 5px;
            padding: 8px 16px;
        }
        QPushButton:hover {
            background-color: #6272a4;
        }
    """)
    
    window = MainWindow()
    window.show()
    
    sys.exit(app.exec_())

if __name__ == "__main__":
    main()
