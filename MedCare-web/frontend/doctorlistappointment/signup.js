function Submitform(formData) {
  console.log("Request sent");
  fetch("/listappointments", {
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
        data.message.sort((a, b) => {
          const dateA = new Date(a.date + ' ' + a.fromdatetime);
          const dateB = new Date(b.date + ' ' + b.fromdatetime);
          return dateA - dateB;
        });

        data.message.forEach(task => {
          const taskCard = document.createElement('div');
          taskCard.classList.add('product');

          taskCard.innerHTML = `
              <div class="basic-details" onclick="DisplayAppoinment(' ${task.appointmentid}',' ${task.customerid}',' ${task.gender}','${task.emailid}','${task.briefdescription}',' ${task.symptoms}','${task.date}',' ${task.fromdatetime}','${task.medications}',' ${task.pastsurgeriestreatments}','${task.notes}','${task.name}','${task.meetlink}')">
              <p>CustomerID: ${task.customerid}</p>
                <span>Name: ${task.name}</span>
                <p>Date of Appointment: ${task.date} ${task.fromdatetime}</p>
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


const tokenDataString = sessionStorage.getItem('doctortoken');
const tokenData = JSON.parse(tokenDataString);
const formData = {
  token: tokenData.token,
};

Submitform(formData);

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
  if (!Object.keys(icon).includes(toastType))
    toastType = "info";

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
function DisplayAppoinment(appointmentid, id, gender, email, dis, symptom, date, time, medications, Past, notes, name, meetlink) {
  document.querySelectorAll('.product').forEach(element => {
    element.style.display = 'none';
  });
  html = document.getElementById('product-container').innerHTML
  document.getElementById('product-container').innerHTML = html + `
    <div class="product1">
    <p onclick="DeleteData()"class="back"><img src="https://www.freepnglogos.com/uploads/x-png/dangerour-x-red-circle-dont-enter-close-health-and-wellness-icon-png-16.png" height="20px"></p>
    <p>Name: ${name}</p>
    <p>AppointmentID: ${appointmentid}</p>
      <p>CustomerID: ${id}</p>
      <p>Gender: ${gender}</p>
      <p>EmailID: ${email}</p>
      <p>Brief Description: ${dis}</p>
      <p>Symptoms: ${symptom}</p>
      <p>Date of Appointment: ${date} ${time}</p>
      <p>Medications: ${medications}</p>
      <p>Past Surgeries/Treatments: ${Past}</p>
      <p>Notes: ${notes}</p>
      <a href="${meetlink}" target="_blank" class="meetlink-button">Meet-Link</a>     
    </div>`;
}
function DeleteData() {
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