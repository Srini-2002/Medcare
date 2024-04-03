function InsertBlog(){
  const formData = {
    title: document.getElementById("title").value,
    content: document.getElementById("content").value,
    author: document.getElementById("author").value,
    pictures: ConvertImageToBase64(document.getElementById("pictures"))
  };

  Promise.all([formData.pictures])
    .then(([photoBase64]) => {
      formData.pictures = photoBase64;
      console.log("Request sent");
      fetch("/insertblog", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      })
        .then(response => response.json())
        .then(data => {
          if (data && data.message) {
           alert(data.message);
          } else {
            alert(data.error);
          }
        })
        .catch(error => {
          // showToast("An error occurred", "danger", 5000);
          alert("An error occurred");
          console.error("Error:", error);
        });
    })
    .catch(error => {
      showToast("An error occurred while processing files", "danger", 5000);
      console.error("Error:", error);
    });
}

function ConvertImageToBase64(imageInput) {
  return new Promise((resolve, reject) => {
    if (imageInput.files.length > 0) {
      const file = imageInput.files[0];
      const reader = new FileReader();
      reader.onload = function (e) {
        const base64String = e.target.result.split(',')[1]; 
        const mimeType = file.type || 'image/png';
        const base64 = 'data:' + mimeType + ';base64,' + base64String;
        resolve(base64);
      };
      reader.readAsDataURL(file); // This line reads the file as a data URL
    } else {
      reject(new Error("Please select an image file"));
    }
  });
}


function DisplayNotApprovedDoctor() {
  document.getElementById('sellersnip').style.display = 'block';
  document.querySelector('.js-blog-cointainer').style.display = 'none'
  document.querySelector('.display-doctor').style.display = 'none';
  let html = ""
  document.querySelector('.seller-list-body').innerHTML = html;
  fetch("/getnotapproveddoctors", {
    method: "POST", 
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify()
  })

    .then(response => response.json())
    .then(data => {



      data.message.forEach(seller => {

        html += `
            <tr class="candidates-list">
            <td class="title">
              <div class="thumb"> <img class="img-fluid"
                  src="${seller.photo}" alt="">
              </div>
              <div class="candidate-list-details">
                <div class="candidate-list-info">
                  <div class="candidate-list-title">
                    <h5 class="mb-0"><a href="#">${seller.name.toUpperCase()}</a></h5>
                  </div>
                  <div class="candidate-list-option">
                    <ul class="list-unstyled">
                      <li><i class="fas fa-filter pr-1"></i>${seller.emailid}
                      </li>
                      <li><i class="fas fa-map-marker-alt pr-1"></i>${seller.specialization}</li>
                    </ul>
                  </div>
                </div>
              </div>
            </td>
            <td class="candidate-list-favourite-time text-center"> <a
                class="candidate-list-favourite order-2 text-danger" href="#"></a>
              <span class="candidate-list-time order-1">${seller.doctorid}</span></td>
            <td>
              <ul class="list-unstyled mb-0 d-flex justify-content-end">
              <li onclick="EditData('${seller.doctorid}','doctor')"><a class="text-info" data-toggle="tooltip" title="" data-original-title="Edit"><i
              class="far fa-eye"></i></a>
              </li>
                <li  onclick="ChangeStatus('${data.message.doctorid}','declined');DisplayListSeller()"><a class="text-danger" data-toggle="tooltip" title=""
                    data-original-title="Delete"><i class="far fa-trash-alt"></i></a></li>
              </ul>
            </td>
          </tr>`;

      });
      document.querySelector('.seller-list-body').innerHTML = html;
    })
    .catch(error => {
      console.log(error)
    });
}


function DisplayApprovedDoctor() {
  document.getElementById('sellersnip').style.display = 'block';
  document.querySelector('.js-blog-cointainer').style.display = 'none'
  document.querySelector('.display-doctor').style.display = 'none';
  let html = ""
  document.querySelector('.seller-list-body').innerHTML = html;
  fetch("/getapproveddoctors", {
    method: "POST", // Use DELETE method to delete data
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify()
  })

    .then(response => response.json())
    .then(data => {
      let html = ""


      data.message.forEach(seller => {

        html += `
            <tr class="candidates-list">
            <td class="title">
              <div class="thumb"> <img class="img-fluid"
                  src="${seller.photo}" alt="">
              </div>
              <div class="candidate-list-details">
                <div class="candidate-list-info">
                  <div class="candidate-list-title">
                    <h5 class="mb-0"><a href="#">${seller.name.toUpperCase()}</a></h5>
                  </div>
                  <div class="candidate-list-option">
                    <ul class="list-unstyled">
                      <li><i class="fas fa-filter pr-1"></i>${seller.emailid}
                      </li>
                      <li><i class="fas fa-map-marker-alt pr-1"></i>${seller.specialization}</li>
                    </ul>
                  </div>
                </div>
              </div>
            </td>
            <td class="candidate-list-favourite-time text-center"> <a
                class="candidate-list-favourite order-2 text-danger" href="#"></a>
              <span class="candidate-list-time order-1">${seller.doctorid}</span></td>
            <td>
              <ul class="list-unstyled mb-0 d-flex justify-content-end">
              <li onclick="ApprovedEditData('${seller.doctorid}','doctor')"><a class="text-info" data-toggle="tooltip" title="" data-original-title="Edit"><i
              class="far fa-eye"></i></a>
              </li>
                <li  onclick="ChangeStatus('${data.message.doctorid}','declined');DisplayListSeller()"><a class="text-danger" data-toggle="tooltip" title=""
                    data-original-title="Delete"><i class="far fa-trash-alt"></i></a></li>
              </ul>
            </td>
          </tr>`;

      });
      document.querySelector('.seller-list-body').innerHTML = html;
    })
    .catch(error => {
      console.log(error)
    });
}



function DisplayDoctor() {
  document.getElementById('sellersnip').style.display = 'block';
  document.querySelector('.js-blog-cointainer').style.display = 'none'
  document.querySelector('.display-doctor').style.display = 'none';
  let html = ""
  document.querySelector('.seller-list-body').innerHTML = html;
  fetch("/getalldoctors", {
    method: "POST", // Use DELETE method to delete data
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify()
  })

    .then(response => response.json())
    .then(data => {

      

      data.message.forEach(seller => {

        html += `
            <tr class="candidates-list">
            <td class="title">
              <div class="thumb"> <img class="img-fluid"
                  src="${seller.photo}" alt="">
              </div>
              <div class="candidate-list-details">
                <div class="candidate-list-info">
                  <div class="candidate-list-title">
                    <h5 class="mb-0"><a href="#">${seller.name.toUpperCase()}</a></h5>
                  </div>
                  <div class="candidate-list-option">
                    <ul class="list-unstyled">
                      <li><i class="fas fa-filter pr-1"></i>${seller.emailid}
                      </li>
                      <li><i class="fas fa-map-marker-alt pr-1"></i>${seller.specialization}</li>
                    </ul>
                  </div>
                </div>
              </div>
            </td>
            <td class="candidate-list-favourite-time text-center"> <a
                class="candidate-list-favourite order-2 text-danger" href="#"></a>
              <span class="candidate-list-time order-1">${seller.doctorid}</span></td>
            <td>
              <ul class="list-unstyled mb-0 d-flex justify-content-end">
              <li onclick="AllEditData('${seller.doctorid}','doctor')"><a class="text-info" data-toggle="tooltip" title="" data-original-title="Edit"><i
              class="far fa-eye"></i></a>
              </li>
                <li  onclick="ChangeStatus('${data.message.doctorid}','declined');DisplayListSeller()"><a class="text-danger" data-toggle="tooltip" title=""
                    data-original-title="Delete"><i class="far fa-trash-alt"></i></a></li>
              </ul>
            </td>
          </tr>`;

      });

      document.querySelector('.seller-list-body').innerHTML = html;
    })
    .catch(error => {
      console.log(error)
    });
}


function EditData(id, profession) {

  fetch("/getdoctor", {
    method: "POST", // Use DELETE method to delete data
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ doctorid: id })
  })

    .then(response => response.json())
    .then(data => {

      let html = ""

      html = `
              <div class="container" style="width:1500px;max-width:1500px; margin-left:300px">
              <div class="row">
              <div class="col-sm-8 col-sm-offset-2">
              <div class="panel panel-white profile-widget">
              <div class="row">
              <div class="col-sm-12">
              <div class="image-container bg2">
              <img src="${data.message.photo}" class="avatar" alt="avatar" height="100px" >
              </div>
              </div>
              <div class="col-sm-12">
              <div class="details">
              <h4>${data.message.name} <i class="fa fa-sheild"></i></h4>
              <h5>${data.message.specialization}</h5>
              <div class="mg-top-10">
              <a onclick="ChangeStatus('${data.message.doctorid}','approved');DisplayListSeller();" class="btn btn-green" style="background-color:green; color:white">Approve</a>
              <a onclick="ChangeStatus('${data.message.doctorid}','declined');DisplayListSeller();" class="btn btn-blue" style="background-color:red; color:white">Decline</a>
              </div>
              </div>
              </div>
              </div>
              </div>
              <div class="row">
              <div class="col-sm-6"   >
              <div class="panel panel-white border-top-purple">
              <div class="panel-heading">
              <h3 class="panel-title" style="height:30px;">Studies</h3>
              </div>
              <div class="panel-body" style="padding:30px; border-radius:5px">
              <div class="body-section">
              <h5 class="section-heading">MedicalSchool : <span class="message">${data.message.medicalschool}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">GraduationDate : <span class="message">${data.message.graduationdate}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">DegreeDocument : <span onclick="Openpdf('${data.message.degreedocument}')"> <a class="text-info" data-toggle="tooltip" title="" data-original-title="Edit"><i
              class="far fa-eye"></i></a> </span> </h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">LicenseNumber <span class="message" >${data.message.licensenumber} </span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">LicenseDocument : <span onclick="Openpdf('${data.message.licensedocument}')"> <a class="text-info" data-toggle="tooltip" title="" data-original-title="Edit"><i
              class="far fa-eye"></i></a> </span> </h5>
              </div>
              <div class="body-section">
              <!-- <a href="#" class="btn btn-purple btn-sm">Edit</a> -->
              </div>
              </div>
              </div>

              <div class="panel panel-white border-top-light-blue">
              <div class="panel-heading">
              <h3 class="panel-title">Account Details</h3>
              </div>
              <div class="panel-body"  style="padding:30px;">
              <div class="body-section">
              <h5 class="section-heading">ID: <span class="message">${data.message.doctorid}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">Created Time: <span class="message" >${data.message.createdtime}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading approved">Approved: <span class="message" >${data.message.isapproved}</span></h5>
              </div>
              </div>
              </div>
              </div>
              <div class="col-sm-6">
              <div class="panel panel-white border-top-green">
              <div class="panel-heading">
              <h3 class="panel-title">User Info</h3>
              <div class="controls pull-right">
              <span class="pull-right clickable">
              <i class="fa fa-chevron-up"></i>
              </span>
              </div>
              </div>
              <div class="panel-body" style="padding:30px">
              <div class="body-section">
              <h5 class="section-heading">Name : <span class="message">${data.message.name}</span></h5>
              </div>

              <div class="body-section">
              <h5 class="section-heading">Telephone:  <span class="message">${data.message.phonenumber}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">Email : <span class="message">${data.message.emailid}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">Date of Birth : <span class="message">${data.message.dateofbirth}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">Gender : <span class="message">${data.message.gender}</span></h5>
              </div>
              </div>
              </div>
              <div class="panel panel-white border-top-orange">
              <div class="panel-heading">
              <h3 class="panel-title">Experience</h3>
              </div>
              <div class="panel-body" style="padding:30px">
              <div class="body-section">
              <h5 class="section-heading">WorkExperience : <span class="message">${data.message.workexperience}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">Specialization : <span class="message" >${data.message.specialization}</span></h5>
              </div>
              </div>
              </div>
    

              
              </div>
              </div>
              </div>
              </div>
              </div>
              </div>
              `;


      document.querySelector('.display-doctor').innerHTML = html;
      document.getElementById('sellersnip').style.display = 'none';
      document.querySelector('.js-blog-cointainer').style.display = 'none'
      document.querySelector('.display-doctor').style.display = 'block';
    })
    .catch(error => {
      console.log(error)
    });

}
function ApprovedEditData(id, profession) {

  fetch("/getdoctor", {
    method: "POST", // Use DELETE method to delete data
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ doctorid: id })
  })

    .then(response => response.json())
    .then(data => {

      let html = ""

      html = `
<div class="container" style="width:1500px;max-width:1500px; margin-left:300px">
              <div class="row">
              <div class="col-sm-8 col-sm-offset-2">
              <div class="panel panel-white profile-widget">
              <div class="row">
              <div class="col-sm-12">
              <div class="image-container bg2">
              <img src="${data.message.photo}" class="avatar" alt="avatar" height="100px" >
              </div>
              </div>
              <div class="col-sm-12">
              <div class="details">
              <h4>${data.message.name} <i class="fa fa-sheild"></i></h4>
              <h5>${data.message.specialization}</h5>
              <div class="mg-top-10">
              <a onclick="ChangeStatus('${data.message.doctorid}','declined');DisplayListSeller();" class="btn btn-blue" style="background-color:red; color:white">Delete</a>
              </div>
              </div>
              </div>
              </div>
              </div>
              <div class="row">
              <div class="col-sm-6"   >
              <div class="panel panel-white border-top-purple">
              <div class="panel-heading">
              <h3 class="panel-title" style="height:30px;">Studies</h3>
              </div>
              <div class="panel-body" style="padding:30px; border-radius:5px">
              <div class="body-section">
              <h5 class="section-heading">MedicalSchool : <span class="message">${data.message.medicalschool}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">GraduationDate : <span class="message">${data.message.graduationdate}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">DegreeDocument : <span onclick="Openpdf('${data.message.degreedocument}')"> <a class="text-info" data-toggle="tooltip" title="" data-original-title="Edit"><i
              class="far fa-eye"></i></a> </span> </h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">LicenseNumber <span class="message" >${data.message.licensenumber} </span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">LicenseDocument : <span onclick="Openpdf('${data.message.licensedocument}')"> <a class="text-info" data-toggle="tooltip" title="" data-original-title="Edit"><i
              class="far fa-eye"></i></a> </span> </h5>
              </div>
              <div class="body-section">
              <!-- <a href="#" class="btn btn-purple btn-sm">Edit</a> -->
              </div>
              </div>
              </div>

              <div class="panel panel-white border-top-light-blue">
              <div class="panel-heading">
              <h3 class="panel-title">Account Details</h3>
              </div>
              <div class="panel-body"  style="padding:30px;">
              <div class="body-section">
              <h5 class="section-heading">ID: <span class="message">${data.message.doctorid}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">Created Time: <span class="message" >${data.message.createdtime}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">Approved: <span class="message" >${data.message.isapproved}</span></h5>
              </div>
              </div>
              </div>
              </div>
              <div class="col-sm-6">
              <div class="panel panel-white border-top-green">
              <div class="panel-heading">
              <h3 class="panel-title">User Info</h3>
              <div class="controls pull-right">
              <span class="pull-right clickable">
              <i class="fa fa-chevron-up"></i>
              </span>
              </div>
              </div>
              <div class="panel-body" style="padding:30px">
              <div class="body-section">
              <h5 class="section-heading">Name : <span class="message">${data.message.name}</span></h5>
              </div>

              <div class="body-section">
              <h5 class="section-heading">Telephone:  <span class="message">${data.message.phonenumber}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">Email : <span class="message">${data.message.emailid}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">Date of Birth : <span class="message">${data.message.dateofbirth}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">Gender : <span class="message">${data.message.gender}</span></h5>
              </div>
              </div>
              </div>
              <div class="panel panel-white border-top-orange">
              <div class="panel-heading">
              <h3 class="panel-title">Experience</h3>
              </div>
              <div class="panel-body" style="padding:30px">
              <div class="body-section">
              <h5 class="section-heading">WorkExperience : <span class="message">${data.message.workexperience}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">Specialization : <span class="message" >${data.message.specialization}</span></h5>
              </div>
              </div>
              </div>
    

              
              </div>
              </div>
              </div>
              </div>
              </div>
              </div>
              `;


      document.querySelector('.display-doctor').innerHTML = html;
      document.getElementById('sellersnip').style.display = 'none';
      document.querySelector('.js-blog-cointainer').style.display = 'none'
      document.querySelector('.display-doctor').style.display = 'block';
    })
    .catch(error => {
      console.log(error)
    });

}
function AllEditData(id, profession) {

  fetch("/getdoctor", {
    method: "POST", // Use DELETE method to delete data
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ doctorid: id })
  })

    .then(response => response.json())
    .then(data => {

      let html = ""

      html = `
<div class="container" style="width:1500px;max-width:1500px; margin-left:300px">
              <div class="row">
              <div class="col-sm-8 col-sm-offset-2">
              <div class="panel panel-white profile-widget">
              <div class="row">
              <div class="col-sm-12">
              <div class="image-container bg2">
              <img src="${data.message.photo}" class="avatar" alt="avatar" height="100px" >
              </div>
              </div>
              <div class="col-sm-12">
              <div class="details">
              <h4>${data.message.name} <i class="fa fa-sheild"></i></h4>
              <h5>${data.message.specialization}</h5>
              <div class="mg-top-10">
              </div>
              </div>
              </div>
              </div>
              </div>
              <div class="row">
              <div class="col-sm-6"   >
              <div class="panel panel-white border-top-purple">
              <div class="panel-heading">
              <h3 class="panel-title" style="height:30px;">Studies</h3>
              </div>
              <div class="panel-body" style="padding:30px; border-radius:5px">
              <div class="body-section">
              <h5 class="section-heading">MedicalSchool : <span class="message">${data.message.medicalschool}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">GraduationDate : <span class="message">${data.message.graduationdate}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">DegreeDocument : <span onclick="Openpdf('${data.message.degreedocument}')"> <a class="text-info" data-toggle="tooltip" title="" data-original-title="Edit"><i
              class="far fa-eye"></i></a> </span> </h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">LicenseNumber <span class="message" >${data.message.licensenumber} </span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">LicenseDocument : <span onclick="Openpdf('${data.message.licensedocument}')"> <a class="text-info" data-toggle="tooltip" title="" data-original-title="Edit"><i
              class="far fa-eye"></i></a> </span> </h5>
              </div>
              <div class="body-section">
              <!-- <a href="#" class="btn btn-purple btn-sm">Edit</a> -->
              </div>
              </div>
              </div>

              <div class="panel panel-white border-top-light-blue">
              <div class="panel-heading">
              <h3 class="panel-title">Account Details</h3>
              </div>
              <div class="panel-body"  style="padding:30px;">
              <div class="body-section">
              <h5 class="section-heading">ID: <span class="message">${data.message.doctorid}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">Created Time: <span class="message" >${data.message.createdtime}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading ">Approved: <span class="message" >${data.message.isapproved}</span></h5>
              </div>
              </div>
              </div>
              </div>
              <div class="col-sm-6">
              <div class="panel panel-white border-top-green">
              <div class="panel-heading">
              <h3 class="panel-title">User Info</h3>
              <div class="controls pull-right">
              <span class="pull-right clickable">
              <i class="fa fa-chevron-up"></i>
              </span>
              </div>
              </div>
              <div class="panel-body" style="padding:30px">
              <div class="body-section">
              <h5 class="section-heading">Name : <span class="message">${data.message.name}</span></h5>
              </div>

              <div class="body-section">
              <h5 class="section-heading">Telephone:  <span class="message">${data.message.phonenumber}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">Email : <span class="message">${data.message.emailid}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">Date of Birth : <span class="message">${data.message.dateofbirth}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">Gender : <span class="message">${data.message.gender}</span></h5>
              </div>
              </div>
              </div>
              <div class="panel panel-white border-top-orange">
              <div class="panel-heading">
              <h3 class="panel-title">Experience</h3>
              </div>
              <div class="panel-body" style="padding:30px">
              <div class="body-section">
              <h5 class="section-heading">WorkExperience : <span class="message">${data.message.workexperience}</span></h5>
              </div>
              <div class="body-section">
              <h5 class="section-heading">Specialization : <span class="message" >${data.message.specialization}</span></h5>
              </div>
              </div>
              </div>
    

              
              </div>
              </div>
              </div>
              </div>
              </div>
              </div>
              `;


      document.querySelector('.display-doctor').innerHTML = html;
      document.getElementById('sellersnip').style.display = 'none';
      document.querySelector('.js-blog-cointainer').style.display = 'none'
      document.querySelector('.display-doctor').style.display = 'block';
    })
    .catch(error => {
      console.log(error)
    });

}

function Openpdf(base64Pdf) {
  const dataUrl = `${base64Pdf}`;
  const newTab = window.open();
  newTab.document.write('<iframe width="100%" height="100%" src="' + dataUrl + '"></iframe>');
};


function ChangeStatus(id,status){
  document.querySelector('.mg-top-10').style.display = 'none';
  if(status == 'approved'){
    document.querySelector('.approved').innerHTML = `Approved: <span class="message" >true</span>`
  }
    const statusData = {
      doctorid:id,
      status:status
  }
  fetch('/changestatus', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(statusData)
    })
        .then(response => response.json())
        .then(data => {
           
            alert('Doctor deleted successfully.');
            
        })
        .catch(error => {
            console.error('Error creating seller:', error);
        });
}

function searchDoctors() {
  const searchInput = document.getElementById('doctorSearch').value.toLowerCase();
  const doctorRows = document.querySelectorAll('.candidates-list');

  doctorRows.forEach(row => {
    const doctorName = row.querySelector('.candidate-list-title h5 a').innerText.toLowerCase();
    if (doctorName.includes(searchInput)) {
      row.style.display = '';
    } else {
      row.style.display = 'none';
    }
  });
}

function DisplayBlog(){
  document.getElementById('sellersnip').style.display = 'none';
  document.querySelector('.display-doctor').style.display = 'none';
  document.querySelector('.js-blog-cointainer').style.display = 'block'
}
