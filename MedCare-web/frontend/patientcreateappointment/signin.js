function Submitform(formdata) {
    console.log("Request sent");
    fetch("/bookappointment", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(formdata),
    })
        .then(response => response.json())
        .then(data => {
            if (data.message) {
                showToast(data.message, "success", 5000); // Fixed typo in toast type
                windowWait('/patienthome'); // Fixed typo in function name
            } else {
                showToast(data.error, "danger", 5000);
            }
            // Token store in session storage

        })
        .catch(error => {
            console.error("Error:", error);
            showToast(error.message, "danger", 5000); // Fixed reference to data.message
        });
}

const tokenDataString = sessionStorage.getItem('token');
const tokenData = JSON.parse(tokenDataString);

document.getElementById("js-signin-button").addEventListener("click", (event) => {
    event.preventDefault();
    
    const formData = {
        token: tokenData.token,
        customerid: "",
        name: document.getElementById("name").value,
        dateofbirth: document.getElementById("date_of_birth").value,
        gender: document.getElementById("gender").value,
        phonenumber: document.getElementById("phone_number").value,
        emailid: document.getElementById("email").value,
        address: document.getElementById("address").value,
        briefdescription: document.getElementById("brief_description").value,
        symptoms: document.getElementById("symptoms").value,
        date: document.getElementById("datetime").value,
        existingconditions: document.getElementById("existing_conditions").value,
        medications: document.getElementById("medications").value,
        pastsurgeriestreatments: document.getElementById("past_surgeries_treatments").value,
        emergencycontactname: document.getElementById("emergency_contact_name").value,
        emergencycontactphonenumber: document.getElementById("emergency_contact_phone_number").value,
        notes: document.getElementById("notes").value
    };

    if (formData.emailid.trim() === "" || formData.name.trim() === "" || formData.dateofbirth.trim() === "" || formData.gender.trim() === "" || formData.phonenumber.trim() === ""
        || formData.address.trim() === "" || formData.briefdescription.trim() === "" || formData.date.trim() === "" || formData.emergencycontactphonenumber.trim() === "") {
        showToast("Please fill in all the fields", "danger", 5000);
        return false;
    }

    Submitform(formData);
});

function windowWait(str) {
    setTimeout(() => {
        window.location.href = str;
    }, 3000);
}

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

    document.body.appendChild(box);
};

function signOut() {
    // Clear session storage
    sessionStorage.clear();
    // Redirect to the sign-in page (replace 'signin.html' with the actual sign-in page URL)
    window.location.href = '/home';
  }
  // Attach the sign-out function to the button click event
  document.getElementById('signout-button').addEventListener('click', signOut);
