function Submitform(formdata) {
    console.log("Request sent");
    fetch("/createprescription", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(formdata),
    })
        .then(response => response.json())
        .then(data => {
            if (data.message) {
                showToast(data.message, "success", 5000);
                WindowWait('/doctorhome')
            } else {
                showToast(data.error, "danger", 5000);
            }
            

        })
        .catch(error => {
            console.error("Error:", error);
            showToast(data.error, "danger", 5000);
        });
}

const tokenDataString = sessionStorage.getItem('doctortoken');
const tokenData = JSON.parse(tokenDataString);

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
function AddMedication() {
    console.log("Called");
    let html = `<div class="form-group">
                    <label for="medication">Medication:</label>
                    <input class="css-input" type="text" name="medication" required>
                    <select class="css-input" name="medication-option" required>
                        <option value="">Select</option>
                        <option value="morning">Morning</option>
                        <option value="afternoon">Afternoon</option>
                        <option value="night">Night</option>
                    </select>
                </div>`;

    document.getElementById("medication-input").insertAdjacentHTML('beforeend', html);
}


function SendMedication(){
    let medicationdata = GetMedicationData()
    const formData = {
        token:tokenData.token,
        appointmentid: document.getElementById("appointmentid").value,
        instruction: document.getElementById("instructions").value,
        nextappointment: document.getElementById("next_appointment").value,
        medication: medicationdata,

    };
    Submitform(formData)
}




function GetMedicationData() {
    console.log("called medication")
    let medications = []; // Array to store medication data

    // Get all medication input elements
    let medicationInputs = document.querySelectorAll('input[name="medication"]');
    let medicationOptions = document.querySelectorAll('select[name="medication-option"]');

    // Iterate over each medication input and option
    for (let i = 0; i < medicationInputs.length; i++) {
        let medication = medicationInputs[i].value; // Get medication value
        let time = medicationOptions[i].value; // Get selected option value (time)

        // Check if medication value is not empty
        if (medication.trim() !== '') {
            // Push medication data into the array
            medications.push({ name: medication, timerange: time });
        }
    }
    return medications; // Return array of medication data
}




