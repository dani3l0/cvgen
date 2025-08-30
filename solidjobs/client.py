import asyncio
import aiohttp
import json
import re
from datetime import datetime


dehtml = re.compile('<.*?>')
def replacer(string):
	string = re.sub(dehtml, "\n", string).replace(".", ". ").replace("\n\n", "\n").replace("&nbsp;", " ")
	return string


async def getAll():
	sess = aiohttp.ClientSession()
	sess.headers.add("Accept", "application/vnd.solidjobs.jobofferlist+json")
	sess.headers.add("Content-Type", "application/vnd.solidjobs.jobofferlist+json")
	resp = await sess.get('https://solid.jobs/api/offers?division=it&sortOrder=default')

	data = await resp.json()
	objs = []
	for job in data:
		# We don't want ones older than 30 days
		added = datetime.fromisoformat(job["validFrom"])
		if (datetime.now().replace(tzinfo=None) - added.replace(tzinfo=None)).days > 30:
			continue
		obj = {}
		obj["id"]				= job["id"]
		obj["title"]			= job["jobTitle"]
		obj["category"]			= job["mainCategory"]
		obj["experienceLevel"]	= job["experienceLevel"]
		obj["experienceYears"]	= job["minimalExperienceInField"] if "minimalExperienceInField" in job else "Nie podano"
		obj["salaryFrom"]		= job["salaryRange"]["lowerBound"]
		obj["salaryTo"]			= job["salaryRange"]["upperBound"]
		obj["salaryPeriod"]		= job["salaryRange"]["salaryPeriod"]
		obj["employmentType"]	= job["salaryRange"]["employmentType"]
		obj["workload"]			= job["workload"]
		obj["company"]			= job["companyName"]
		obj["address"]			= job["companyAddress"]
		obj["city"]				= job["companyCity"]
		obj["validFrom"]		= job["validFrom"]
		obj["validTo"]			= job["validTo"]
		obj["logo"]				= job["companyLogoUrl"]
		obj["languages"]		= [x["name"] for x in job["requiredLanguages"]]
		obj["skills"]			= [x["name"] for x in job["requiredSkills"]]
		obj["remote"]			= job["remotePossible"]
		obj["map"]				= f"https://www.openstreetmap.org/search?lat={job["officeLatitude"]}&lon={job["officeLongitude"]}&zoom=9#map=9/{job["officeLatitude"]}/{job["officeLongitude"]}"
		obj["url"]				= f"https://solid.jobs/offer/{obj["id"]}/{job['jobOfferUrl']}"
		obj["apiUrl"]			= f"https://solid.jobs/api/offers/{obj["id"]}/{job['jobOfferUrl']}"
		objs.append(obj)
	await sess.close()
	return objs


async def getMore(obj):
	sess = aiohttp.ClientSession()
	sess.headers.add("Accept", "application/vnd.solidjobs.jobofferdetails+json")
	sess.headers.add("Content-Type", "application/vnd.solidjobs.jobofferdetails+json")
	resp2 = await sess.get(obj["apiUrl"])
	data2 = await resp2.json()
	await sess.close()
	data2 = data2["jobOfferDetails"]
	obj["dutiesDescription"] = replacer(data2["jobDescription"])
	obj["requirements"] = replacer(data2["candidateProfile"])
	return obj
