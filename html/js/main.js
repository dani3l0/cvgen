// Config files
var CONFIG, CONTENTS

// Rendering function
var RENDER

// HTTP request wrapper
var reqId = 0
function req(url, callback) {
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

// Load config & contents
document.addEventListener("DOMContentLoaded", () => {
	req("config.json", conf => {
		CONFIG = JSON.parse(conf)
		document.querySelector("#ifr").src = `templates/${CONFIG.template_name}/index.html`
		req(CONFIG.contents_json, contents => {
			CONTENTS = JSON.parse(contents)
			req(`templates/${CONFIG.template_name}/contents.js`, engine => {
				eval(engine)
				if (typeof(RENDER) == "function") {
					console.log("Render engine loaded, rendering data...")
					RENDER(document.querySelector("#ifr").contentDocument, CONTENTS)
				} else {
					console.log("Failed to load rendering engine")
				}
			})
		})
	})
})
