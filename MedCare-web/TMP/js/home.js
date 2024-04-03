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


    

