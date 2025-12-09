// Config files
var CONFIG, CONTENTS

// Rendering function
var RENDER

// HTTP request wrapper
var reqId = 0
const req = (url, callback) => {
	let req = reqId++
	console.log(`Request #${req}: pending   [${url}]`)
	let xhr = new XMLHttpRequest()
	xhr.open("GET", url)
	xhr.onload = function() {
		console.log(`Request #${req}: done      [${url}]`)
		callback(this.responseText)
	}
	xhr.send()
}

// Text replacer; replaces fake variables with real text
var fillContent = () => {
	const d = document.querySelector("#ifr").contentDocument
	for (let key in CONTENTS) {
		let e = CONTENTS[key]
		if (typeof(e) == "string") e = [e]
		for (let i = 0; i < e.length; i++) {
			d.body.innerHTML = d.body.innerHTML.replace(`%{${key}}`, e[i])
		}
	}
}

// Load config & contents
document.addEventListener("DOMContentLoaded", () => {
	req("config.json", conf => {
		CONFIG = JSON.parse(conf)
		document.querySelector("#ifr").src = `templates/${CONFIG.template_name}/index.html`
		req(CONFIG.contents_json, contents => {
			CONTENTS = JSON.parse(contents)
			req(`templates/${CONFIG.template_name}/render.js`, engine => {
				eval(engine)
				document.querySelector("#ifr").addEventListener("load", () => {
					if (typeof(RENDER) == "function") {
						console.log("Render engine loaded, rendering data...")
						RENDER(document.querySelector("#ifr").contentDocument, CONTENTS)
						fillContent()
					} else {
						console.log("Failed to load rendering engine")
					}
				})
			})
		})
	})
})
