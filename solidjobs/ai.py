from os import path
import ollama


# Some globals
globalPrompt = "Automatyzuję proces szukania pracy i masz za zadanie analizować mi treść ogłoszenia. Przeanalizuj podane dane i odpowiedz jednym słowem 'tak' lub 'nie', a w następnej linijce krótko uzasadnij swoją decyzję."
nono = "Możesz przymknąć oko na drobne rzeczy, bo praktycznie nigdy nie da się spełnić wszystkich wymagań z ogłoszenia."
ai = ollama.Client(host="127.0.0.1")

# Descriptions
def loadDescs():
	print("\nLoading candidate description file")
	global AboutMe
	AboutMe = open(path.join("../data/AboutMe.txt"), "r").read().strip()


# Generic chat function
def answer(message : str, noglobalprompt=False) -> str:
	message = f"{globalPrompt}\n\n{message}" if not noglobalprompt else message
	m = [{'role': 'user', 'content': message}]
	return ai.chat(model="gpt-oss:20b", messages=m).message.content


# Checks if position is ok
def position(job : dict) -> str:
	if "php" in job["title"].lower():
		print("kurwa php")
		return None
	for x in range(2):
		prompt = "Czy to stanowisko mogłoby w jakikolwiek sposób pasować do moich umiejętności? Jeśli cokolwiek może mieć ono wspólnego, zaakceptuj je.\n"
		prompt += f"- Nazwa stanowiska z ogłoszenia: {job["title"]}\n"
		prompt += f"## Moje umiejętności:\n{AboutMe}"
		resp = answer(prompt)
		if _yes(resp):
			desc = _explain(resp)
			print(f"Brzmi dobrze: {desc}")
			return desc
	print("To raczej nie to")
	return None


# Checks if experience is ok
def experience(job : dict) -> str:
	prompt = "Czy spełniam wymóg lat doświadczenia zawodowego (koniecznie zawodowego)?\n"
	prompt += f"- Wymagane lata doświadczenia: {job["experienceYears"]} ({job["experienceLevel"]})\n"
	prompt += f"## Moje umiejętności (patrz tylko na doświadczenie zawodowe):\n{AboutMe}"
	prompt += "Jeśli nie podano wymagań ad. doświadczenia, po prostu je zaakceptuj."
	resp = answer(prompt)
	if _yes(resp):
		desc = _explain(resp)
		print(f"Chyba jest git: {desc}")
		return desc
	print(f"jprdl, musi być {job["experienceYears"]} lat doświadczenia 🤦‍")
	return None


# Checks whether job requirements are fulfilled
def requirements(job : dict) -> str:
	for x in range(2):
		prompt = f"Czy spełniam wymagania z ogłoszenia? {nono}\n\n"
		prompt += f"## Wymagania:\n{job["requirements"]}\n\n"
		prompt += f"## Moje umiejętności:\n{AboutMe}"
		resp = answer(prompt)
		desc = _explain(resp)
		if _yes(resp):
			print(f"No nie wierzę, spełniasz wymagania: {desc}")
			return desc
	print(f"kogoś chyba cycki szczypią: {desc}")
	return None


# Checks whether duties are appropriate to skills
def duties(job : dict) -> str:
	for x in range(2):
		prompt = f"Czy poradziłbym sobie w wykonywaniu obowiązków podanych w ogłoszeniu? {nono}\n\n"
		prompt += f"## Obowiązki:\n{job["dutiesDescription"]}\n\n"
		prompt += f"## Moje umiejętności:\n{AboutMe}"
		resp = answer(prompt)
		desc = _explain(resp)
		if _yes(resp):
			print(f"Nie powinno być za dużo do roboty: {desc}")
			return desc
	print(f"ale zapierdol: {desc}")
	return None


# Other shit
def _yes(text : str) -> bool:
	y = "tak" in text.split("\n")[0].lower()
	return y and _explain(text)
def _explain(text : str) -> str:
	splot = text.split("\n", 2)
	if len(splot) < 2:
		return None
	return splot[1].strip()
