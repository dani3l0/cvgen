import smtplib, ssl
import random
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
from email.mime.image import MIMEImage
import time
from datetime import datetime


##############################################
smtp_server = "smtp.my.email.domain"
port = 587
sender_email = "aiworkfinder@my.email.domain"
password = "fuccHRbitches"
to_addr = "my@email.address"
base_url = "http://127.0.0.1:56765"
##############################################

def niceDate(date : str) -> str:
	parsd = datetime.fromisoformat(date)
	return parsd.strftime("%A, %d %B %Y")


def sendMessage(job : dict):
	server = smtplib.SMTP(smtp_server, port)
	server.starttls(context=ssl.create_default_context())
	server.login(sender_email, password)

	# Parse markdown text
	hellos = ["Czołem, nierobie!", "Cześć huju!", "Dziń dybry!", "Witam klasę niepracującą,", "Czołem bałwanie,"]
	message = f"""
		<h2>{random.choice(hellos)}</h2>
		Coś ci znalazłem cymbale jeden ty:
		<ul>
			<li><i>Stanowisko</i>: <b>{job['title']}</b></li>
			<li><i>Firma</i>: <b>{job['company']}</b></li>
			<li><i>Miasto</i>: <b>{job['city']}</b></li>
			<li><i>Adres</i>: <b>{job['address']}</b></li>
			<li><i>Zarobki</i>: <b>od {job['salaryFrom']} do {job['salaryTo']} PLN ({job['salaryPeriod']})</b></li>
			<li><i>Papiery</i>: <b>{job['employmentType']}</b></li>
			<li><i>Etat</i>: <b>{job['workload']}</b></li>
			<li><i>Możliwe zdalne</i>: <b>{job['remote']}</b></li>
			<li><i>Umiejętności</i>: <b>{', '.join(job['skills'])}</b></li>
		</ul>
		<i>Dodano {niceDate(job['validFrom'])}</i>
		<i>Ważne do {niceDate(job['validTo'])}</i>
		<h3><a href="{job['url']}">Link do zasranego ogłoszenia</a></h3>
		<h3><a href="{base_url}/{job["id"]}/index.html">Link do wygenerowanego CV</a></h3>
	""".strip().replace("\n", "<br>")

	# Build image object
	msg = MIMEMultipart()
	msg["Subject"] = "Solid.Dżobs: znalazłem jakiegoś Januszexa"
	msg["From"] = f"AI Work Finder <{sender_email}>"
	msg["To"] = to_addr

	# Insert image
	image = MIMEImage(open("../data/screenshot.png", "rb").read())
	image.add_header("Content-ID", f"<{job["id"]}>")
	msg.attach(image)
	message += f"""<br><img src="cid:{job["id"]}">"""

	# Prepare text contents
	msg.attach(MIMEText(message, "html"))
	server.sendmail(msg["From"], [msg["To"]], msg.as_string())
