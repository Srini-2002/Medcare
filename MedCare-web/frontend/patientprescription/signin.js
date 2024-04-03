function SubmitForm(formData) {
  console.log("Request sent");
  fetch("/listpatientprescription", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(formData),
  })
    .then(response => response.json())
    .then(data => {
      const productContainer = document.getElementById('product-container');
      if (data && data.message) {
        showToast("Fetched", "success", 5000);
        data.message.sort((a, b) => new Date(b.createdtime) - new Date(a.createdtime));

        data.message.forEach(task => {
          const taskCard = document.createElement('div');
          taskCard.classList.add('product');

          taskCard.addEventListener('click', function() {
            DisplayAppointment(task.customerid, task.appointmentid, task.doctorid, task.instruction, task.nextappointment, task.medication, task.createdtime);
           });
          taskCard.innerHTML = `
              <div class="basic-details" onclick="DisplayAppointment('${task.customerid}', '${task.appointmentid}', '${task.doctorid}', '${task.instruction}', '${task.nextappointment}', ${task.medication}, '${task.createdtime}')">
                <p>AppointmentID: ${task.appointmentid}</p>
                <span>DoctorID: ${task.doctorid}</span>
                <p>NextAppointment: ${task.nextappointment}</p>
              </div>
          `;
          productContainer.appendChild(taskCard);
        });
      } else {
        showToast(data.error, "danger", 5000);
      }
    })
    .catch(error => {
      showToast(error.message, "danger", 5000);
    });
}

const tokenDataString = sessionStorage.getItem('token');
const tokenData = JSON.parse(tokenDataString);
const formData = {
  token: tokenData.token,
};

SubmitForm(formData);

const icon = {
  success: '<span class="material-symbols-outlined">task_alt</span>',
  danger: '<span class="material-symbols-outlined">error</span>',
  warning: '<span class="material-symbols-outlined">warning</span>',
  info: '<span class="material-symbols-outlined">info</span>',
};

const showToast = (
  message = "Sample Message",
  toastType = "info",
  duration = 5000) => {
  toastType = icon[toastType] ? toastType : "info";

  const box = document.createElement("div");
  box.classList.add("toast", `toast-${toastType}`);
  box.innerHTML = `
    <div class="toast-content-wrapper">
      <div class="toast-message">${message}</div>
      <div class="toast-progress"></div>
    </div>`;
  duration = duration || 5000;
  box.querySelector(".toast-progress").style.animationDuration =
    `${duration / 1000}s`;

  const toastAlready = document.body.querySelector(".toast");
  if (toastAlready) {
    toastAlready.remove();
  }

  document.body.appendChild(box);
};

function DisplayAppointment(customerid, appointmentid, doctorid, instruction, nextappointment, medication, createdtime) {
  document.querySelectorAll('.product').forEach(element => {
    element.style.display = 'none';
  });

  const productContainer = document.getElementById('product-container');
  const backButton = `
    <p onclick="DeleteData()" class="back">
      <img src="https://www.freepnglogos.com/uploads/x-png/dangerour-x-red-circle-dont-enter-close-health-and-wellness-icon-png-16.png" height="20px">
    </p>`;

  let medicationHTML = '';
  if (Array.isArray(medication)) {
    medicationHTML = medication.map(med => `<p>Name:${med.name}</p>
    <p>TimeRange: ${med.timerange}</p>`).join('');
  } else {
    medicationHTML = '<p>No medication information available</p>';
  }

  productContainer.innerHTML += `
    <div class="product1">
      ${backButton}
      <p>CustomerID: ${customerid}</p>
      <p>AppointmentID: ${appointmentid}</p>
      <p>DoctorID: ${doctorid}</p>
      <p>Instruction: ${instruction}</p>
      <p>Next Appointment: ${nextappointment}</p>
      <div class="medication-info">
        <p>Medication:</p>
        ${medicationHTML}
      </div>
    </div>`;
}

function DeleteData(){
  document.querySelectorAll('.product1').forEach(element => {
    element.remove();
  });
  document.querySelectorAll('.product').forEach(element => {
    element.style.display = 'block';
  });
}

function WindowWait(str) {
  setTimeout(() => {
    window.location.href = str;
  }, 3000);
}

function signOut() {
  sessionStorage.clear();
  window.location.href = '/home';
}

// Attach the sign-out function to the button click event
document.getElementById('signout-button').addEventListener('click', signOut);
