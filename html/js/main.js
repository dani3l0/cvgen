var CONFIG, CONTENTS

// Load config & contents
document.addEventListener("DOMContentLoaded", e => {
	let xhr = new XMLHttpRequest()
	xhr.open("GET", "config.json")
	xhr.onload = function() {
		CONFIG = JSON.parse(this.responseText)
		let xhr = new XMLHttpRequest()
		document.querySelector("#ifr").src = `templates/${CONFIG.template_name}/index.html`
		xhr.open("GET", CONFIG.contents_json)
		xhr.onload = function() {
			CONTENTS = JSON.parse(this.responseText)
			fill()
		}
		xhr.send()
	}
	xhr.send()
})

// Replace contents text, fillup CV with data
function fill() {
	let e = document.querySelector("#ifr").contentDocument
}
