let menu = document.querySelector('#menu-btn');
let navbar = document.querySelector('.navbar');

menu.onclick = () =>{
    menu.classList.toggle('fa-times');
    navbar.classList.toggle('active');
}

window.onscroll = () =>{
    menu.classList.remove('fa-times');
    navbar.classList.remove('active');
}

document.querySelector('.book-appointment-button').addEventListener('click',(event)=>{
    var messageElement = document.querySelector(".book-appointment-button");

    messageElement.innerHTML = "Please  Sign In to Book an Appointment";

    const id = setTimeout(()=>{
        var messageElement = document.querySelector(".book-appointment-button");
        messageElement.innerHTML = "Book Appoinment";
        event.preventDefault()
    },3500)
})





function signOut() {
    // Clear session storage
    sessionStorage.clear();
    // Redirect to the sign-in page (replace 'signin.html' with the actual sign-in page URL)
    window.location.href = '/signin';
}
// Attach the sign-out function to the button click event
document.getElementById('signout-button').addEventListener('click', signOut);


    

