function Submitform(formData) {
  console.log("Request sent");
  fetch("/listpatientreport", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(formData),
  })
    .then(response => response.json())
    .then(data => {
      if (data && data.message) {
        // Display the medical report
        showMedicalReport(data.message);
        showToast("Fetched", "success", 5000);
      } else {
        showToast(data.error, "danger", 5000);
      }
    })
    .catch(error => {
      showToast("Error fetching data", "danger", 5000);
    });
};

function showMedicalReport(data) {
  // Check if the report container already exists
  let reportContainer = document.querySelector('.medical-report-container');
  if (!reportContainer) {
    // If not, create it
    reportContainer = document.createElement('div');
    reportContainer.classList.add('medical-report-container');
    document.body.appendChild(reportContainer);
  }

  // Create a new container for the report
  let report = document.createElement('div');
  report.classList.add('container');
  report.innerHTML = `
    <!DOCTYPE html>
    <html lang="en">
    <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Medical Report</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 20px;
        }
        .container1 {
          width: 800px;
          padding: 20px;
          font-size: larger;

          display: inline-block;
          margin-top: 1rem;
          padding: .5rem;
          padding-left: 1rem;
          border: var(--border);
          border-radius: .5rem;
          box-shadow: var(--box-shadow);
          color: #595c5c;
          cursor: pointer;
          font-size: 1.7rem;
          background: #fff;
          margin-left: 500px;
      }
      
        h1 {
            text-align: center;
        }
        .section {
            margin-bottom: 20px;
        }
        .section-title {
            font-weight: bold;
            margin-bottom: 10px;
        }
        .section-content {
            margin-left: 20px;
        }
        .print-button {
          display: inline-block;
          margin-top: 1rem;
          padding: .5rem;
          padding-left: 1rem;
          border: var(--border);
          border-radius: .5rem;
          box-shadow: var(--box-shadow);
          color: var(--green);
          cursor: pointer;
          font-size: 1.7rem;
          background: #fff;
          margin-left: 50px;
            
        }
        .print-button:hover {
          background: var(--green);
          color: #fff;
      }

    
      .print-button {
        display: inline-block;
        margin-top: 1rem;
        padding: .5rem;
        padding-left: 1rem;
        border: var(--border);
        border-radius: .5rem;
        box-shadow: var(--box-shadow);
        color: var(--green);
        cursor: pointer;
        font-size: 1.7rem;
        background: #fff;
        margin-left: 300px;
    }
    
    .print-button button {
        padding: .7rem 1rem;
        border-radius: .5rem;
        background: var(--green);
        color: #fff;
        margin-left: .5rem;
    }
    
    .print-button button:hover {
        background: var(--green);
        color: #fff;
    }
    
    .print-button button:hover span {
        color: var(--green);
        background: #fff;
        margin-left: 1rem;
    }
    @media print {
      body {
          margin: 0;
          padding: 0;
      }
      .container1 {
          margin: auto;
          width: 80%; /* Adjust the width as needed */
          padding: 20px;
          font-size: larger;
          display: block; /* Ensure it's a block element */
          margin-top: 1rem;
          padding: .5rem;
          padding-left: 1rem;
          border: var(--border);
          border-radius: .5rem;
          box-shadow: var(--box-shadow);
          color: #595c5c;
          cursor: pointer;
          font-size: 1.7rem;
          background: #fff;
          margin-left: 20px; /* Reset margin-left */
      }
  }
    
    
    </style>
    </head>
    <body>
    <div class="container1">
        <h1>Medical Report</h1>
        <div class="section">
            <div class="section-title">Patient Information</div>
            <div class="section-content">
                <p><strong>Name:</strong> ${data.name}</p>
                <p><strong>Date of Birth:</strong> ${data.dateofbirth}</p>
                <p><strong>Gender:</strong> ${data.gender}</p>
                <p><strong>Phone Number:</strong> ${data.phonenumber}</p>
                <p><strong>Email:</strong> ${data.emailid}</p>
                <p><strong>Address:</strong> ${data.address}</p>
            </div>
        </div>
        <div class="section">
            <div class="section-title">Appointment Details</div>
            <div class="section-content">
                <p><strong>Symptoms:</strong> ${data.symptoms}</p>
                <p><strong>Date:</strong> ${data.date}</p>
                <p><strong>Time:</strong> ${data.fromdatetime} - ${data.todatetime}</p>
                <p><strong>Doctor Specialization:</strong> ${data.doctorspecialization}</p>
                <p><strong>Preferred Doctor:</strong> ${data.preferreddoctorid}</p>
            </div>
        </div>
        <div class="section">
            <div class="section-title">Prescriptions</div>
            <div class="section-content">
                ${data.prescriptions.map(prescription => {
    return `
                    <p><strong>Instruction:</strong> ${prescription.instruction}</p>
                    <p><strong>Next Appointment:</strong> ${prescription.nextappointment}</p>
                    <p><strong>Medications:</strong></p>
                    <ul>
                      ${prescription.medication.map(med => {
      return `<li>${med.name} - ${med.timerange}</li>`;
    }).join('')}
                    </ul>
                  `;
  }).join('')}
            </div>
        </div>
        <div class="print-button">
            <button onclick="window.print()">Print Report</button>
        </div>
    </div>
    </body>
    </html>
  `;

  // Append the report to the report container
  reportContainer.innerHTML = '';
  reportContainer.appendChild(report);
}





const tokenDataString = sessionStorage.getItem('token');
const tokenData = JSON.parse(tokenDataString);

function ListRecord() {
  console.log("-----------------")
  const formData = {
    token: tokenData.token,
  };
  Submitform(formData)
}


ListRecord()


let icon = {
  success:
    '<span class="material-symbols-outlined">task_alt</span>',
  danger:
    '<span class="material-symbols-outlined">error</span>',
  warning:
    '<span class="material-symbols-outlined">warning</span>',
  info:
    '<span class="material-symbols-outlined">info</span>',
};

const showToast = (
  message = "Sample Message",
  toastType = "info",
  duration = 5000) => {
  if (
    !Object.keys(icon).includes(toastType))
    toastType = "info";

  let box = document.createElement("div");
  box.classList.add(
    "toast", `toast-${toastType}`);
  box.innerHTML = ` <div class="toast-content-wrapper"> 
                      <div class="toast-message">${message}</div> 
                      <div class="toast-progress"></div> 
                      </div>`;
  duration = duration || 5000;
  box.querySelector(".toast-progress").style.animationDuration =
    `${duration / 1000}s`;

  let toastAlready =
    document.body.querySelector(".toast");
  if (toastAlready) {
    toastAlready.remove();
  }

  document.body.appendChild(box)
};

function WindowWait(str) {
  setTimeout(() => {
    window.location.href = str
  }, 3000)

}

function signOut() {
  sessionStorage.clear();
  window.location.href = '/home';
}
// Attach the sign-out function to the button click event
document.getElementById('signout-button').addEventListener('click', signOut);