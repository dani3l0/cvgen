RENDER = (doc, contents) => {
	const v = (text) => {return `%{${text}}`}

	// Main
	doc.querySelector(".first-name").innerText = v("first_name")
	doc.querySelector(".last-name").innerText = v("last_name")
	doc.querySelector(".specialization-summary").innerText = v("summary")
	doc.querySelector(".policy i").innerText = v("policy_note_icon")
	doc.querySelector(".policy div").innerText = v("policy_note")

	// Contact
	let contact = doc.querySelector(".contact").cloneNode(true)
	let contact_item = contact.querySelector(".item")
	contact.querySelector(".name").innerText = v("contact_title")
	for (let item of contact.querySelectorAll(".item")) contact.removeChild(item)
	for (let _ of contents.contact_icons) {
		let i = contact_item.cloneNode(true)
		i.querySelector("i").innerText = v("contact_icons")
		if (_ == "github") {
			let gh = document.createElement("img")
			gh.src = "img/github.svg"
			i.classList.add("github")
			i.querySelector("i").replaceWith(gh)
		}
		i.querySelector(".value").innerText = v("contact_values")
		i.querySelector(".desc").innerText = v("contact_descs")
		contact.appendChild(i)
	}
	doc.querySelector(".contact").innerHTML = contact.innerHTML

	// Education
	let edu = doc.querySelector(".education").cloneNode(true)
	let edu_item = edu.querySelector(".item")
	edu.querySelector(".name").innerText = v("education_title")
	for (let item of edu.querySelectorAll(".item")) edu.removeChild(item)
	for (let _ of contents.education_values) {
		let i = edu_item.cloneNode(true)
		i.querySelector("i").innerText = v("education_icons")
		i.querySelector(".value div").innerText = v("education_values")
		i.querySelector(".desc").innerText = v("education_descs")
		edu.appendChild(i)
	}
	doc.querySelector(".education").innerHTML = edu.innerHTML

	// Experience
	let exp = doc.querySelector(".experience").cloneNode(true)
	let exp_item = exp.querySelector(".item")
	exp.querySelector(".name").innerText = v("experience_title")
	for (let item of exp.querySelectorAll(".item")) exp.removeChild(item)
	for (let _ of contents.experience_values) {
		let i = exp_item.cloneNode(true)
		i.querySelector("i").innerText = v("experience_icons")
		i.querySelector(".value div").innerText = v("experience_values")
		i.querySelector(".desc").innerText = v("experience_descs")
		exp.appendChild(i)
	}
	doc.querySelector(".experience").innerHTML = exp.innerHTML

	// About me
	doc.querySelector(".about-me .name").innerText = v("about_me_title")
	doc.querySelector(".about-me .text").innerText = v("about_me_note")

	// Skills
	doc.querySelector(".skills .name").innerText = v("skills_title")
	doc.querySelector(".skills .desc").innerText = v("skills_note")
	let box = doc.querySelector(".skillbox").cloneNode(true)
	let skill_item = box.querySelector(".item")
	for (let item of box.querySelectorAll(".item")) box.removeChild(item)
	let j = 0
	let colors = contents.skills_colors
	for (let _ of contents.skills_values) {
		let i = skill_item.cloneNode(true)
		i.className = ""
		i.setAttribute("style", `--color: ${colors[j]}`)
		i.classList.add("item", `b${++j}`)
		i.querySelector("i").innerText = v("skills_icons")
		i.querySelector(".title div").innerText = v("skills_values")
		i.querySelector(".text").innerText = v("skills_descs")
		let bubbles = i.querySelector(".bubbles")
		for (let item of bubbles.querySelectorAll("div")) bubbles.removeChild(item)
		for (let _ of contents[`skills_bubbles_${j}`]) {
			let b = document.createElement("div")
			b.innerText = `%{skills_bubbles_${j}}`
			i.querySelector(".bubbles").appendChild(b)
		}
		console.log(i)
		box.appendChild(i)
	}
	doc.querySelector(".skillbox").innerHTML = box.innerHTML

	// Achievements
	doc.querySelector(".achievements .name").innerText = v("achievements_title")
	doc.querySelector(".achievements .text").innerText = v("achievements_note")
	let achievs = doc.querySelector(".competition").cloneNode(true)
	let achiev = achievs.querySelector(".item")
	achievs.innerHTML = ""
	const places = ["gold", "silver", "bronze", "none"]
	for (let place of contents.achievements_places) {
		let i = achiev.cloneNode(true)
		i.classList.remove("gold", "silver", "bronze")
		i.classList.add(places[Math.min(place - 1, 3)])
		i.querySelector(".title").innerText = v("achievements_values")
		i.querySelector(".desc").innerText = v("achievements_descs")
		achievs.appendChild(i)
	}
	doc.querySelector(".competition").innerHTML = achievs.innerHTML
}
