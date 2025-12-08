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

	try {
		let req = await fetch(`/api/${apipath}`, { headers, method, body })
		status = req.status
		json = await req.json()
	} catch (e) {
		console.error("Failed parsing request json: ", e)
	}

	let end = new Date().getTime()
	return {
		ok: status == 200,
		status: status,
		json: json,
		respTimeMs: Math.ceil(end - start)
	}
}