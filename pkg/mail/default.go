package mail

// Default :
type Default struct{}

// Name :
func (hm *Default) Name() string {
	return "default"
}

// HTMLTemplate :
func (hm *Default) HTMLTemplate() string {
	return `
	<html>
	<head>
		<link href="https://fonts.googleapis.com/css?family=Nunito" rel="stylesheet" type="text/css">
		<style>
			body {
				font-family: Nunito;
				font-size: 14px;
				-webkit-text-size-adjust: 100%;
				-ms-text-size-adjust: 100%;
				margin: 0 !important;
				padding: 0;
				/* background-color: #f0f0f0; */
			}
			
			table {
				border-spacing: 0;
				font-family: sans-serif;
				color: #333333;
			}
			
			td {
				padding: 0;
			}
			
			img {
				border: 0;
			}
			
			div[style*="margin: 16px 0"] {
				margin: 0 !important;
			}
			
			.wrapper {
				width: 100%;
				table-layout: fixed;
				-webkit-text-size-adjust: 100%;
				-ms-text-size-adjust: 100%;
			}
			
			.webkit {
				max-width: 600px;
				margin: 0 auto;
			}
			
			.outer {
				Margin: 0 auto;
				width: 100%;
				max-width: 600px;
			}
			
			.full-width-image img {
				width: 100%;
				max-width: 100px;
				height: auto;
			}
			
			.inner {
				padding: 10px;
			}
			
			p {
				Margin: 0;
			}
			
			a {
				color: #ee6a56;
				text-decoration: underline;
			}

			.no-decor {
				text-decoration: none;
			}
			
			.h1 {
				font-size: 21px;
				font-weight: bold;
				Margin-bottom: 18px;
			}
			
			.h2 {
				font-size: 18px;
				font-weight: bold;
				Margin-bottom: 12px;
			}
			/* One column layout */
			
			.one-column .contents {
				text-align: left;
			}
			
			.one-column p {
				font-size: 14px;
				Margin-bottom: 10px;
			}
			/*Two column layout*/
			
			.two-column {
				text-align: center;
				font-size: 0;
			}
			
			.two-column .column {
				width: 100%;
				max-width: 300px;
				display: inline-block;
				vertical-align: top;
			}
			
			.contents {
				width: 100%;
			}
			
			.two-column .contents {
				font-size: 14px;
				text-align: left;
			}
			
			.two-column img {
				width: 100%;
				max-width: 280px;
				height: auto;
			}
			
			.two-column .text {
				padding-top: 10px;
			}
			/*Three column layout*/
			
			.three-column {
				text-align: center;
				font-size: 0;
				padding-top: 10px;
				padding-bottom: 10px;
			}
			
			.three-column .column {
				width: 100%;
				max-width: 200px;
				display: inline-block;
				vertical-align: top;
			}
			
			.three-column.store .column {
				width: 100%;
				max-width: 50px;
				display: inline-block;
				vertical-align: top;
			}
			
			.three-column .contents {
				font-size: 14px;
				text-align: center;
			}
			
			.three-column.social-media img {
				width: 100%;
				max-width: 20px;
				height: auto;
				margin: 5px;
			}
			
			.three-column .text {
				padding-top: 10px;
			}
			/* Left sidebar layout */
			
			.left-sidebar {
				text-align: center;
				font-size: 0;
			}
			
			.left-sidebar .column {
				width: 100%;
				display: inline-block;
				vertical-align: middle;
			}
			
			.left-sidebar .left {
				max-width: 100px;
			}
			
			.left-sidebar .right {
				max-width: 250px;
			}
			
			.left-sidebar .img {
				width: 100%;
				max-width: 80px;
				height: auto;
			}
			
			.left-sidebar .contents {
				font-size: 14px;
				text-align: center;
			}
			
			.left-sidebar a {
				color: #85ab70;
			}

			.md-container p{
				text-align: left;
				margin-left: 20px;
				margin-right: 20px;
				font-size: 16px;
				line-height: 25px;
				color: #575657;
				margin-bottom: 25px;
			}

		</style>
	</head>
	
	<body>
		<div style="background-color: #f0f0f0;">
			<table class="outer" align="center">
				<tr>
					<td class="left-sidebar">
						<div class="column right" style="max-width: 100%">
							<table width="100%">
								<tr>
									<td class="inner">
										<p style="margin-top: 50px; margin-bottom: 20px; text-align: center;">
											<img src="https://gruperdev.s3-ap-southeast-1.amazonaws.com/Assets/Blue-Lines.jpg" width="300" alt="" />
										</p>
									</td>
								</tr>
							</table>
						</div>
					</td>
				</tr>
	
				<tr style="background-color: #fff;">
					<td class="one-column">
						<table width="100%">
							<tr>
								<!-- MAIN CONTENT -->
								<!-- Hubungi Kami -->
								{{ if (eq .Mail.MailType "contactUs" ) }}								
								<td class="inner contents">
									<p style="text-align: left; margin-left: 20px; margin-right: 20px; font-size: 20px; margin-top: 40px; margin-bottom: 40px">
										Halo <span style="font-weight: bold">Admin</span>,<br>
									</p>
									<p style="text-align: left; margin-left: 20px; margin-right: 20px; font-size: 16px; line-height: 25px; color: #575657; margin-bottom: 25px">
										Anda mendapatkan pesan yang dikirim oleh pengguna melalui form pada halaman "Hubungi Kami" di website Supergroup.
									</p>
									<ul style="list-style:none;font-size: 16px;">
										<li style="margin-bottom: 10px;">Nama Pengirim : <span style="font-weight: bold">{{ .Format.ContactUs.Name }}</span></li>
										<li style="margin-bottom: 10px;">Email Pengirim : <span style="font-weight: bold">{{ .Format.ContactUs.Email }}</li>
										<li style="margin-bottom: 10px;">Judul : <span style="font-weight: bold">{{ .Format.ContactUs.Subject }}</li>
										<li style="margin-bottom: 10px;">Pesan : <span style="font-weight: bold">{{ .Format.ContactUs.Message }}</li>
									</ul>
									<p style="text-align: left; margin-left: 20px; margin-right: 20px; font-size: 16px; line-height: 25px; color: #575657; margin-bottom: 25px">
										Jika Anda ingin membalas pesan tersebut, silahkan mengirimkan email balasan Anda ke alamat email pengirim yang tertera diatas.
									</p>
									<p style="text-align: left; margin-left: 20px; margin-right: 20px; font-size: 16px; line-height: 25px; color: #575657; margin-bottom: 25px">
										Terima kasih
									</p>
								</td>
								{{ end }}

								<!-- Forgot -->
								{{ if (eq .Mail.MailType "forgot" ) }}
								<td class="inner contents">
									<p style="text-align: left; margin-left: 20px; margin-right: 20px; font-size: 20px; margin-top: 40px; margin-bottom: 40px">Halo <span style="font-weight: bold">{{ .Format.Forgot.Name }}</span>,</p>
									<p style="text-align: left; margin-left: 20px; margin-right: 20px; font-size: 16px; line-height: 25px; color: #575657; margin-bottom: 50px">
										Kami menerima permintaan untuk mengganti password Anda. Untuk melakukan pergantian password, silahkan menekan tombol dibawah ini :
									</p>
									<p style="text-align: center; margin-left: 20px; margin-right: 20px; font-size: 14px; line-height: 25px; color: #575657; margin-bottom: 50px;">
										<a href="{{ .Format.Forgot.ButtonLink }}" style="color: #fff; text-decoration: none; padding-top: 15px; padding-bottom: 15px; padding-left: 40px; padding-right: 40px; background-color: #1771f8; border-radius: 3px;">Ganti Password</a>
									</p>
									<p style="text-align: left; margin-left: 20px; margin-right: 20px; font-size: 16px; line-height: 25px; color: #575657; margin-bottom: 25px">
										Link ini akan aktif hingga 60 menit sejak email ini terkirim. Jika bukan Anda yang meminta untuk mengganti password, maka silahkan abaikan email ini.
									</p>
								</td>
								{{ end }}

								<!-- ENDOF MAIN CONTENT -->
							</tr>
						</table>
					</td>
				</tr>
	
				<tr style="background-color: #fcfcfc">
					<td class="one-column">
						<table width="100%">
							<tr>
								<td class="inner contents">
									<p class="h1" style="text-align: center; font-size: 16px
									; margin-top: 60px">Gratis. <span style="font-weight: normal">Download App sekarang</span></p>
								</td>
							</tr>
						</table>
					</td>
				</tr>
	
				<tr style="background-color: #fcfcfc">
					<td class="three-column store">
						<div class="column">
							<table width="100%">
								<tr>
									<td class="inner">
										<table class="contents">
											<tr>
												<td class="text"></td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
						</div>
						<div class="column" style="max-width: 400px">
							<table width="100%">
								<tr>
									<td class="inner">
										<table class="contents">
											<tr>
												<td class="text">
													<p style="margin-top: -35px; margin-bottom: 60px">
														<a href="https://www.apple.com/ios/app-store"><img src="https://gruperdev.s3-ap-southeast-1.amazonaws.com/Assets/app+store.png" width="150" alt="" /></a>
														<a href="https://play.google.com/store"><img src="https://gruperdev.s3-ap-southeast-1.amazonaws.com/Assets/play+store.png" width="150" alt="" /></a>
													</p>
												</td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
						</div>
						<div class="column">
							<table width="100%">
								<tr>
									<td class="inner">
										<table class="contents">
											<tr>
												<td class="text"></td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
						</div>
					</td>
				</tr>
	
				<tr style="background-color: #fff;">
					<td class="one-column">
						<table width="100%">
							<tr>
								<td class="inner contents">
									<p style="text-align: center; font-size: 16px; margin-top: 25px;">Punya pertanyaan?</p>
									<p style="text-align: center; font-size: 16px; margin-bottom: 25px;"><a href="mailto:kusnandartoni@gmail.com" style="color: #1771f8">kusnandartoni@gmail.com</a></p>
								</td>
							</tr>
						</table>
					</td>
				</tr>
	
				<tr>
					<td class="three-column social-media">
						<div class="column">
							<table width="100%">
								<tr>
									<td class="inner">
										<table class="contents">
											<tr>
												<td class="text"></td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
						</div>
						<div class="column">
							<table width="100%">
								<tr>
									<td class="inner">
										<table class="contents">
											<tr>
												<td class="text">
													<p>
														<a href="https://www.linkedin.com/in/kusnandartoni/" target="_blank" class="no-decor">
															<img src="https://gruperdev.s3-ap-southeast-1.amazonaws.com/Assets/ig-logo.png" width="20" height="auto" alt="" />
														</a>
														<a href="https://www.linkedin.com/in/kusnandartoni/" target="_blank" class="no-decor">
															<img src="https://gruperdev.s3-ap-southeast-1.amazonaws.com/Assets/tw-logo.png" width="20" height="auto" alt="" />
														</a>
														<a href="https://www.linkedin.com/in/kusnandartoni/" target="_blank" class="no-decor">
															<img src="https://gruperdev.s3-ap-southeast-1.amazonaws.com/Assets/fb-logo.png" width="20" height="auto" alt="" />
														</a>
													</p>
												</td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
						</div>
						<div class="column">
							<table width="100%">
								<tr>
									<td class="inner">
										<table class="contents">
											<tr>
												<td class="text"></td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
						</div>
					</td>
				</tr>
	
				<tr>
					<td class="one-column">
						<table width="100%">
							<tr>
								<td class="inner contents">
									<p style="text-align: center; color: #999999;">Copyrights 2019. Kusnandartoni</p>
									<p style="text-align: center; color: #999999;">Jakarta</p>
									<p style="text-align: center; color: #999999; margin-bottom: 50px">Indonesia</p>
								</td>
							</tr>
						</table>
					</td>
				</tr>
			</table>
		</div>
	</body>
	<html>

	`
}

// PlainTextTemplate :
func (hm *Default) PlainTextTemplate() string {
	return ``
}
