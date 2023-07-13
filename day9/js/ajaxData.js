// GET data api
const promises = new Promise((ada, gada) => {
    const xhr = new XMLHttpRequest()

    // open, onload, onerror, send
    xhr.open("GET", "https://api.npoint.io/ef0bd8024f1eb74d12aa", true)
    xhr.onload = () => {
        if(xhr.status === 200) {
            ada(JSON.parse(xhr.response))
        } else if (xhr.status >= 400) {
            gada("Error loading data")
        }
    }
    xhr.onerror = () => {
        gada("Jaringan Bermasalah . . .")
    }
    xhr.send()
})
// buat variabel kosong
let dataTestimonials = []

const datas = async (rating) => {
    try {
        const res = await promises
        dataTestimonials = res
        console.log( res)
        allTestimonial()
    } catch (error) {
        console.error("error hehe")
    }
}
datas()

function allTestimonial() {
    // alert("ok")
    let testimonualHtml = ""
    dataTestimonials.forEach((d) => {
        testimonualHtml += `<div class="cardTesti">
                 <img alt="p" src="${d.img}"/>
                 <h5>${d.quote}</h5>
                 <p>${d.user}  <i class="fa-solid fa-star"></i> ${d.rating} </p>
             </div>
         `
    })
    document.getElementById("boxTesti").innerHTML = testimonualHtml

}
allTestimonial()


function filterData(rating) {
    let testimonialshtmlfilter = ""

    const datafilter = dataTestimonials.filter((d) => {
        return d.rating === rating
    })

    datafilter.forEach((d) => {
        testimonialshtmlfilter += `<div class="cardTesti">
                <img alt="p" src="${d.img}"/>
                <h5>${d.quote}</h5>
                <p>${d.user}  <i class="fa-solid fa-star"></i> ${d.rating} </p>
            </div>
        `
    })


    document.getElementById("boxTesti").innerHTML = testimonialshtmlfilter

}

function cari() {
    let search = document.getElementById("search").value.toLowerCase();
    // console.log(search)
    testimonialshtmlcari = ""

    const dataCari = dataTestimonials.filter((d) => {
        // mengonversi kata kunci pencarian menjadi huruf kecil
        // cari nilai dalam array
        return d.user.toLowerCase().includes(search);
    })
    dataCari.forEach((d) => {
        testimonialshtmlcari += `<div class="cardTesti">
                <img alt="p" src="${d.img}"/>
                <h5>${d.quote}</h5>
                <p>${d.user}  <i class="fa-solid fa-star"></i> ${d.rating} </p>
            </div>
        `
    })
   
    document.getElementById("boxTesti").innerHTML = testimonialshtmlcari

}
