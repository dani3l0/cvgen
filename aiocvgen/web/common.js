var fetch2 = async (apipath, postdata) => {
	let start = new Date().getTime()
	let body
	let method = "GET"
	if (postdata) {
		method = "POST"
		body = JSON.stringify(postdata)
	}

	let headers = new Headers()
	if (postdata) headers.append("Content-Type", "application/json")
	let status = 0
	let json = {}
	let text = ""

	try {
		let req = await fetch(`/api/${apipath}`, { headers, method, body })
		status = req.status
		text = await req.text()
		json = JSON.parse(text)
	} catch (e) {
		json = text
	}

	let end = new Date().getTime()
	return {
		ok: status == 200,
		status: status,
		json: json,
		respTimeMs: Math.ceil(end - start)
	}
}

let notifTimeout
var notif = (text) => {
	let div = document.querySelector("#status")
	div.innerHTML = text
	clearTimeout(notifTimeout)
	notifTimeout = setTimeout(() => div.innerHTML = "", 4000)
}
