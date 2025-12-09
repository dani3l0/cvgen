import asyncio
import ai
import client
import mail
import pyppeteer
import os
import base64
import json
import webserver
from datetime import datetime, timedelta


async def main():
	# Start AI & Test AI
	print("\nStarting AI")
	print(f"Available models: {[x.model for x in ai.ai.list().models]}")
	print("\nTesting AI")
	print(ai.answer("Cześć, powiedz króciutko w jednym zdaniu że działasz poprawnie", noglobalprompt=True))

	# Read about me information
	ai.loadDescs()

	# Run webserver
	asyncio.create_task(webserver.serve_httpd())

	# Read already parsed jobs so we can skip them
	print("\nReading already parsed jobs")
	lastjob = None
	try:
		checked = json.loads(open("./data/checkedJobs.json", "r").read())
		print(f"Found {len(checked)} checked jobs, those will be skipped")
	except FileNotFoundError:
		checked = []
		print("No parsed jobs file found. Assuming there are no checked jobs yet")

	# Main l00p
	while True:
		for job in (await client.getAll()):

			# Check if already parsed
			if lastjob and lastjob["id"] not in checked:
				checked.append(lastjob["id"])
				w = open("./data/checkedJobs.json", "w+")
				w.write(json.dumps(checked))
				w.close()
			elif job["id"] in checked:
				continue
			lastjob = job

			# Analyzing started
			print(f"\nAnalizuję: '{job["title"]}'")

			# Check if position is the one I wanted
			position = ai.position(job)
			if not position:
				continue
		
			# Check required experience
			exp = ai.experience(job)
			if not exp:
				continue

			# Get more info for further steps
			job = await client.getMore(job)

			# Check requirements
			reqs = ai.requirements(job)
			if not reqs:
				continue
		
			# Check duties
			duties = ai.duties(job)
			if not duties:
				continue

			# Seems like this job is for me
			print("OOO KURWA")
		
			# Prepare fucking CV
			while True:
				try:
					jsoned = ai.answer(f"""
						Przygotuj jak najbardziej dopasowaną treść idealnego kandydata w formacie json do poniższego tekstu z ogłoszenia:

						## Treść z ogłoszenia:
						{job["requirements"]}
						{job["dutiesDescription"]}

						## Wymagany format json i opis danych:
						{{
							"about_me_note": "Krótki, idealny opis kandydata w 1. osobie. Charakter, styl pracy, chęć nauki itd. Od 70 do 90 słów",
							"skills_note": "Krótkie podsumowanie wszystkich umiejętności, takie chwytliwe hasło; Od 24 do 40 słów",
							"skills_values": ["Tytuł pierwszej sekcji umiejętności, tylko jedno słowo", "Tytuł drugiej sekcji umiejętności, tylko jedno słowo", "Tytuł trzeciej sekcji umiejętności, tylko jedno słowo", "Tytuł czwartej sekcji umiejętności, tylko jedno słowo"],
							"skills_descs": ["Opis umiejętności z pierwszej sekcji; Od 20 do 28 słów", "Opis umiejętności z drugiej sekcji; Od 20 do 28 słów", "Opis umiejętności z trzeciej sekcji; Od 20 do 28 słów", "Opis umiejętności z czwartej sekcji; Od 20 do 28 słów"],
							"skills_bubbles_1": ["Piewszy tag/słowo klucz do pierwszej sekcji", "Drugi tag/słowo klucz do pierwszej sekcji", "3 do 5 tagów"],
							"skills_bubbles_2": ["Piewszy tag/słowo klucz do drugiej sekcji", "Drugi tag/słowo klucz do drugiej sekcji", "3 do 5 tagów"],
							"skills_bubbles_3": ["Piewszy tag/słowo klucz do trzeciej sekcji", "Drugi tag/słowo klucz do trzeciej sekcji", "3 do 5 tagów"],
							"skills_bubbles_4": ["Piewszy tag/słowo klucz do czwartej sekcji", "Drugi tag/słowo klucz do czwartej sekcji", "3 do 5 tagów"],
							"skills_icons": ["Ikonka pasująca do tytułu pierwszej sekcji; użyj wartości zgodnej z MaterialIcons", "Ikonka pasująca do tytułu drugiej sekcji; użyj wartości zgodnej z MaterialIcons", "Ikonka pasująca do tytułu trzeciej sekcji; użyj wartości zgodnej z MaterialIcons", "Ikonka pasująca do tytułu czwartej sekcji; użyj wartości zgodnej z MaterialIcons"]
						}}

						Wypisz TYLKO I WYŁĄCZNIE wygenerowany json, który można przeparsować. Na dodatek, wszystkie wartości mają być w języku polskim.
					""", noglobalprompt=True)
					jsoned = jsoned.replace("```json", "").replace("```", "")
					print(jsoned)
					jsoned = json.loads(jsoned)
					break
				except:
					print(jsoned)
					continue
			pth = f"./data/CVs/{job["id"]}"
			os.system(f"cp -r './html' {pth}") # veeery unsafe :P

			# Modify json contents file for customized CV
			orig = json.loads(open(pth+"/contents.json", "r").read())
			for key in jsoned:
				orig[key] = jsoned[key]
			with open(pth+"/contents.json", "w") as f:
				f.write(json.dumps(orig, indent=4))
				f.close()

			# Take a screenshot showing where the Januszex is
			print("Launching browser")
			browser = await pyppeteer.launch({"defaultViewport": {"width": 960, "height": 540}})
			page = await browser.newPage()
			print("Browser launched")
			print("Taking map screenshot")
			await page.goto(job["map"])
			await asyncio.sleep(3)
			await page.evaluate("""() => {
				let target = "#map .leaflet-pane .leaflet-marker-pane .leaflet-marker-icon:nth-child(2)";
				let node = document.querySelector(target);
				node.style.display = "none";
			}""")
			await page.screenshot({"path": "./data/screenshot.png"})
			print("Screenshot taken")
			print("Closing browser")
			await browser.close()

			# Send message & cleanup
			mail.sendMessage(job)
			os.remove("./data/screenshot.png")

		await asyncio.sleep(3600)


asyncio.run(main())
