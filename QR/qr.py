import qrcode


local_ip_url = "http://www.tomato.com" 


qr = qrcode.QRCode(
    version=1,
    error_correction=qrcode.constants.ERROR_CORRECT_L,
    box_size=10,
    border=4,
)

qr.add_data(local_ip_url)
qr.make(fit=True)


img = qr.make_image(fill='black', back_color='white')


img.save("local_ip_3001_qr.png")
