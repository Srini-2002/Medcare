
function Submitform(formData) {


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
        showToast("Fetched", "success", 5000);
      } else {
        showToast(data.error, "danger", 5000);
      }

      WindowWait()
    })
    .catch(error => {
      showToast(data.error, "danger", 5000);
    });
};


const tokenDataString = sessionStorage.getItem('token');
const tokenData = JSON.parse(tokenDataString);

function ListRecord(){

  const formData = {
    token : tokenData.token,
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