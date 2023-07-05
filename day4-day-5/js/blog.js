//buat array kosong
let blogr = [];



function tambahBlog(event) {
    //untuk menghilangkan default reload

    event.preventDefault()
    //masukan ionc kedalam variabel
    let nodejsic = ' <i class="fa-brands fa-node"></i>';
    let reactic = '<i class="fa-brands fa-react"></i>';
    let typeScriptic = '<i class="fa-brands fa-square-js"></i>';
    let nextic = '<i class="fa-brands fa-vuejs"></i>';

    //masukan elemen berdasarkan id dengan value yg di input kedalam variabel
    let nama = document.getElementById("nama").value;
    let start = document.getElementById("start").value;
    let endd = document.getElementById("endDate").value;
    let deskripsi = document.getElementById("deskripsi").value;
    //masukan element berdasarkan id dengan cheklis yang di pilih kedalam variabel
    let iconnode = document.getElementById("nodeJss").checked ? nodejsic : "";
    let iconreact = document.getElementById("react").checked ? reactic : "";
    let iconnext = document.getElementById("next").checked ? nextic : "";
    let typeScript = document.getElementById("typeScript").checked ? typeScriptic : "";
    let image = document.getElementById("image").files;

    //mencari value kalkulasi
    let x = new Date(start)
    let y = new Date(endd)
    let tanggalSekarang = x.getTime();
    let tanggalDurasi = y.getTime();
    let jarak = tanggalDurasi - tanggalSekarang;

    // mengubah format dari imput kalender
    const listBulan = ['jan', 'feb', 'mar', 'apr', 'may', 'jun', 'jul', 'aug', 'sep', 'oct', 'nov', 'dec'];
    // tanggal
    let dateTimestart = x.getDay()
    let dateTimeend = y.getDay()
    // Bulan
    let dateStart = listBulan[x.getMonth()]
    let dateEnd = listBulan[y.getMonth()]
    // Tahun
    let dateTahunStart = x.getFullYear()
    let dateTahunend = y.getFullYear()

    // menghitung durasi atau mengkalkulasi
    let detik = Math.floor(jarak / 1000);       // detik
    let menit = Math.floor(detik / 60);         //detik
    let jam = Math.floor(menit / 60);           //jam
    let hari = Math.floor(jam / 24);            //tanggal
    let bulan = Math.floor(hari / 30);          //bulan
    let tahun = Math.floor(bulan / 12);         //tahun

  

    //upload image ke URL
    image = URL.createObjectURL(image[0]);
    console.log(image);

    //buat object
    let blogs = {
        nama,

        // detik,
        // menit,
        // jam,
        hari,
        bulan,
        tahun,
        jarak,
        // tanggalDurasi,
        // tanggalSekarang,

        dateTimestart,
        dateTimeend,
        dateStart,
        dateEnd,
        dateTahunStart,
        dateTahunend,

        start,
        endd,
        deskripsi,
        iconnode,
        iconreact,
        iconnext,
        typeScript,
        image,
    };

    //array kosong bloger masukan/tambahkan object blogs kedalam
    blogr.push(blogs);
    console.log(blogr);

    //memanggil function ketika function tambah blog di jalankan
    viewElement();

}

function viewElement() {
    //ketika di klik
    //document berdasarkan id contentView ubah menjadi kosong
    document.getElementById("view").innerHTML = "";

    for (let index = 0; index < blogr.length; index++) {
        document.getElementById("view").innerHTML += `
        <div class="view-grup">
            <div class="view-img">
                <img src="${blogr[index].image}" alt="badro" >
            </div>
            <div class="veiw-desk">
                <h1>${blogr[index].nama}</h1>
                <p>Dari ${blogr[index].dateTimestart} ${blogr[index].dateStart}  ${blogr[index].dateTahunStart}   </p>
                <p>Sampai ${blogr[index].dateTimeend} ${blogr[index].dateEnd}  ${blogr[index].dateTahunend}   </p>
                <p>Durasi ${blogr[index].hari} hari atau ${blogr[index].bulan} Bulan, ${blogr[index].tahun} tahun</p>
                <p>${blogr[index].deskripsi}</p>
            </div>
            <div class="iconss">
                ${blogr[index].iconnode}
                ${blogr[index].iconreact}
                ${blogr[index].iconnext}
                ${blogr[index].typeScript}
                </div>
            <div class="buttonGrup">
            <button><a href="detail.html">Detail</a> </button>
                <button>Delete</button>
            </div>
        </div>
    `;
    }

}
