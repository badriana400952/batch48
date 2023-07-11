class TestimonialClass {
    #quote = ""
    #img = ""

    constructor(quote, img) {
        this.#quote = quote;
        this.#img = img;
    }

    get quote() {
        return this.#quote
    }
    get img() {
        return this.#img
    }
    get user() {
        throw new Error('there is must be user to make testimonials')
    }

    get dataHatml() {
        return `<div class="cardTesti">
                <img alt="p" src="${this.img}"/>
                <h5>${this.quote}</h5>
                <p>${this.user}  </p>
            </div>
        `
    }

}

class userTestimonial extends TestimonialClass {
    #user = ""

    constructor(user, quote, img) {
        super(quote, img)
        this.#user = user
    }
    get user() {
        return "user : " + this.#user
    }
}

class campeniTestimonial extends TestimonialClass {
    #campeni = ""

    constructor(campeni, quote, img) {
        super(quote, img)
        this.#campeni = campeni
    }
    get user() {
        return "campeni : " + this.#campeni

    }
}

const testimonialsBadri = new userTestimonial("badriana", "slebwee", "https://images.unsplash.com/photo-1685239218538-11d65c8baaa1?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=464&q=80")
const testimonialsBayu = new userTestimonial("bayu", "wleeee", "https://images.unsplash.com/photo-1685239218538-11d65c8baaa1?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=464&q=80")
const testimonialsBahrudin = new campeniTestimonial("bahrudin", "slebwee", "https://images.unsplash.com/photo-1685239218538-11d65c8baaa1?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=464&q=80")

let dataTestimonial = [testimonialsBadri, testimonialsBayu, testimonialsBahrudin]
let dataHatml = ""

for (let i = 0; i < dataTestimonial.length; i++) {
    dataHatml += dataTestimonial[i].dataHatml
}
document.getElementById("boxTesti").innerHTML = dataHatml