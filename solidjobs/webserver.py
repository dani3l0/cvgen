import os
from aiohttp import web

async def serve_httpd():
	print("\nStarting httpd")
	os.makedirs("./data/CVs", exist_ok=True)
	app = web.Application()
	app.add_routes([web.static("/", "./data/CVs", show_index=True)])
	await web._run_app(app, host="0.0.0.0", port=56765)
