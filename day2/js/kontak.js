function handleKirim() {
    let nama  = document.getElementById('nama').value;
    let email = document.getElementById('email').value;
    let telepon = document.getElementById('phone').value;
    let skils = document.getElementById('selec').value;
    let deskripsi = document.getElementById('deskripsi').value;


    let emIlPenerima = 'badriana0402@gmail.com';
let a = document.createElement('a');
a.href = `mailto:${emIlPenerima}?subject=${skils}&body=hallo badriana nama saya ${nama}
%0D%0A Email : ${email}
%0D%0A Phone Number : ${telepon}
%0D%0A Deskripsi : ${deskripsi}
`;
a.click()


}


