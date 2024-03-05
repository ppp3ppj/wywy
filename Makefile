templ: 
	@templ generate -watch -proxy=http://localhost:1323

tailwind:
	@tailwindcss -i ./assets/css/styles.css -o ./public/tailwind.css --watch
