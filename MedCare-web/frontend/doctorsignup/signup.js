function SubmitForm(formData) {
  console.log("Request sent");
  fetch("/signupdoctor", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(formData),
  })
    .then(response => response.json())
    .then(data => {
      console.log("Response received:", data.value);
      if (data && data.value) {
        showToast(data.value, "success", 5000);
        WindowWait('/doctorsignin');
      } else {
        showToast(data.error, "danger", 5000);
      }
    })
    .catch(error => {
      showToast("An error occurred", "danger", 5000);
      console.error("Error:", error);
    });
}

document.getElementById("js-signin-button").addEventListener("click", (event) => {
  event.preventDefault();
  const formData = {
    doctorid: "",
    name: document.getElementById("name").value,
    dateofbirth: document.getElementById("date_of_birth").value,
    gender: document.getElementById("gender").value,
    emailid: document.getElementById("email").value,
    phonenumber: document.getElementById("phone_number").value,
    password: document.getElementById("password").value,
    confirmpassword: document.getElementById("confirm_password").value,
    medicalschool: document.getElementById("medical_school").value,
    graduationdate: document.getElementById("graduation_date").value,
    degreedocument: ConvertPdftoBase64(document.getElementById("degree_document")),
    licensenumber: document.getElementById("license_number").value,
    licensedocument: ConvertPdftoBase64(document.getElementById("license_document")),
    workexperience: document.getElementById("work_experience").value,
    specialization: document.getElementById("specialization").value,
    photo: ConvertImageToBase64(document.getElementById("photo"))
  };
  Promise.all([formData.degreedocument, formData.licensedocument, formData.photo])
    .then(([degreeDocBase64, licenseDocBase64, photoBase64]) => {
      formData.degreedocument = degreeDocBase64;
      formData.licensedocument = licenseDocBase64;
      formData.photo = photoBase64;


      if (formData.name.trim() === '' || formData.emailid.trim() === "" || formData.password.trim() === "" || formData.confirmpassword.trim() === "" ||
        formData.dateofbirth.trim() === '' || formData.gender.trim() === '' || formData.phonenumber.trim() === '' || formData.workexperience.trim() === '' ||
        formData.specialization.trim() === '') {
        showToast("Please fill in all the fields", "danger", 5000);
        return false;
      }

      if (formData.confirmpassword.trim() !== formData.password.trim()) {
        showToast("Password Mismatch", "danger", 5000);
        return false;
      }

      SubmitForm(formData);
    })
    .catch(error => {
      showToast("An error occurred while processing files", "danger", 5000);
      console.error("Error:", error);
    });
});

const icon = {
  success: '<span class="material-symbols-outlined">task_alt</span>',
  danger: '<span class="material-symbols-outlined">error</span>',
  warning: '<span class="material-symbols-outlined">warning</span>',
  info: '<span class="material-symbols-outlined">info</span>',
};

const showToast = (message = "Sample Message", toastType = "info", duration = 5000) => {
  if (!Object.keys(icon).includes(toastType)) {
    toastType = "info";
  }
  const box = document.createElement("div");
  box.classList.add("toast", `toast-${toastType}`);
  box.innerHTML = ` <div class="toast-content-wrapper"> 
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

function WindowWait(str) {
  setTimeout(() => {
    window.location.href = str;
  }, 3000);
}

function ConvertPdftoBase64(pdfInput) {
  return new Promise((resolve, reject) => {
    if (pdfInput.files.length > 0) {
      const file = pdfInput.files[0];
      const reader = new FileReader();
      reader.onload = function (e) {
        const base64String = e.target.result.split(',')[1];
        resolve('data:application/pdf;base64,' + base64String);
      };
      reader.readAsDataURL(file); // This line was missing
    } else {
      reject(new Error("Please select a PDF file"));
    }
  });
}

function ConvertImageToBase64(imageInput) {
  return new Promise((resolve, reject) => {
    if (imageInput.files.length > 0) {
      const file = imageInput.files[0];
      const reader = new FileReader();
      reader.onload = function (e) {
        const base64String = e.target.result.split(',')[1];
        // Assuming the file type is always image, change 'image/png' accordingly if needed
        const mimeType = file.type || 'image/png';
        const base64 = 'data:' + mimeType + ';base64,' + base64String;
        // Do something with the base64 string, like displaying it or sending it to server
        resolve(base64);
      };
      reader.readAsDataURL(file); // This line reads the file as a data URL
    } else {
      reject(new Error("Please select an image file"));
    }
  });
}

function resetFormFields() {
  document.getElementById("name").value = '';
  document.getElementById("email").value = '';
  document.getElementById("password").value = '';
  document.getElementById("confirm_password").value = '';
  document.getElementById("date_of_birth").value = '';
  document.getElementById("gender").value = '';
  document.getElementById("phone_number").value = '';
  document.getElementById("medical_school").value = '';
  document.getElementById("graduation_date").value = '';
  document.getElementById("license_number").value = '';
  document.getElementById("work_experience").value = '';
  document.getElementById("specialization").value = '';
  document.getElementById("photo").value = '';
  document.getElementById("degree_document").value = '';
  document.getElementById("license_document").value = '';
}
