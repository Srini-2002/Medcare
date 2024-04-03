function Submitform(formdata) {
    console.log("Request sent");
    fetch("/signincustomer", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(formdata),
    })
        .then(response => response.json())
        .then(data => {
            if (data.token) {
                showToast("Login Success", "success", 5000);
                const tokenDataString = JSON.stringify(data);
                sessionStorage.setItem('token', tokenDataString);
                WindowWait('/patienthome')
            } else {
                showToast(data.error, "danger", 5000);
            }
            //token store in session storage

        })
        .catch(error => {
            console.error("Error:", error);
            showToast(data.message, "danger", 5000);
        });
}



document.getElementById("js-signin-button").addEventListener("click", (event) => {
    event.preventDefault();
    const formData = {
        emailid: document.getElementById("email").value,
        password: document.getElementById("password").value,
    };

    if (formData.emailid.trim() === "" || formData.password.trim() === "") {
        showToast("Please fill in all the fields", "danger", 5000);
        return false;
    }

    Submitform(formData)
})


function WindowWait(str) {
    setTimeout(() => {
      window.location.href = str
    }, 3000)
  
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

    document.body.appendChild(box)
};