// class TestimonialClass {
//     #quote = ""
//     #img = ""

//     constructor(quote, img) {
//         this.#quote = quote;
//         this.#img = img;
//     }

//     get quote() {
//         return this.#quote
//     }
//     get img() {
//         return this.#img
//     }
//     get user() {
//         throw new Error('there is must be user to make testimonials')
//     }

//     get dataHatml() {
//         return `<div class="cardTesti">
//                 <img alt="p" src="${this.img}"/>
//                 <h5>${this.quote}</h5>
//                 <p>${this.user}  </p>
//             </div>
//         `
//     }

// }

// class userTestimonial extends TestimonialClass {
//     #user = ""

//     constructor(user, quote, img) {
//         super(quote, img)
//         this.#user = user
//     }
//     get user() {
//         return "user : " + this.#user
//     }
// }

// class campeniTestimonial extends TestimonialClass {
//     #campeni = ""

//     constructor(campeni, quote, img) {
//         super(quote, img)
//         this.#campeni = campeni
//     }
//     get user() {
//         return "campeni : " + this.#campeni

//     }
// }

// const testimonialsBadri = new userTestimonial("badriana", "slebwee", "https://images.unsplash.com/photo-1685239218538-11d65c8baaa1?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=464&q=80")
// const testimonialsBayu = new userTestimonial("bayu", "wleeee", "https://images.unsplash.com/photo-1685239218538-11d65c8baaa1?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=464&q=80")
// const testimonialsBahrudin = new campeniTestimonial("bahrudin", "slebwee", "https://images.unsplash.com/photo-1685239218538-11d65c8baaa1?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=464&q=80")

// let dataTestimonial = [testimonialsBadri, testimonialsBayu, testimonialsBahrudin]
// let dataHatml = ""

// for (let i = 0; i < dataTestimonial.length; i++) {
//     dataHatml += dataTestimonial[i].dataHatml
// }
// document.getElementById("boxTesti").innerHTML = dataHatml


const dataTestimonial = [
    {
        user: "Badriana",
        quote: "slebew tetew",
        img: "https://images.unsplash.com/photo-1685239218538-11d65c8baaa1?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=464&q=80",
        rating: 1,

    },
    {
        user: "Nurizati islamiyah",
        quote: "terserah",
        img: "https://images.unsplash.com/photo-1687360441372-757f8b2b6835?ixlib=rb-4.0.3&ixid=M3wxMjA3fDF8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870&q=80",
        rating: 2,

    },
    {
        user: "Bayu",
        quote: "a minta duit",
        img: "https://images.unsplash.com/photo-1516898671633-efce17ccb6e0?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=896&q=80",
        rating: 2,

    },
    {
        user: "Bahrudin",
        quote: "pinjem duit",
        img: "https://images.unsplash.com/photo-1499952127939-9bbf5af6c51c?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=876&q=80",
        rating: 4,

    },
    {
        user: "ikhsan",
        quote: "ngopi wa",
        img: "https://images.unsplash.com/photo-1554151228-14d9def656e4?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=386&q=80",
        rating: 4,

    },
    {
        user: "Beni",
        quote: "slebew sensasional ",
        img: "https://images.unsplash.com/photo-1438761681033-6461ffad8d80?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870&q=80",
        rating: 5,

    },
    {
        user: "Epok",
        quote: "haruh",
        img: "https://images.unsplash.com/photo-1491349174775-aaafddd81942?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80",
        rating: 5,

    },
]

function allTestimonial() {
    // alert("ok")
    let testimonualHtml = ""
    dataTestimonial.forEach((d) => {
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

    const datafilter = dataTestimonial.filter((d) => {
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

    const dataCari = dataTestimonial.filter((d) => {
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
