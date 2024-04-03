import re
from PyPDF2 import PdfReader

# Function to extract the patient's name from the PDF text
def extract_patient_name(pdf_text):
    # Split the text into lines
    lines = pdf_text.split('\n')
    # Loop through each line to find the line containing "Name"
    for line in lines:
        # Check if the line contains "Name"
        if 'Name' in line:
            # Extract the patient's name (assuming it comes after "Name")
            # You may need to adjust the logic here based on the actual format of your PDF
            name = line.split('Name')[1].strip()
            return name
    return None

# Specify the path to your PDF file
pdf_path = "MedReport.pdf"

# Creating a pdf reader object
reader = PdfReader(pdf_path)

# Extracting text from the first page
page_text = reader.pages[0].extract_text()

# Extracting patient's name
patient_name = extract_patient_name(page_text)

# Print patient's name
print("Patient's Name:", patient_name)
