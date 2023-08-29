package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Telephone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#ea5a47" d="M62.298 16.452c-8-6-44-6-52 0c0 0-4 4-2 10h13v-5c6-4 24-4 30 0v5h13c2-6-2-10-2-10z"/><path fill="#d22f27" d="m50.971 21.146l13.876-.088l-.441 5.038l-12.463-.088zm-43.841 0l13.877-.088l-.442 5.038l-12.463-.088z"/><path fill="#ea5a47" d="M24.875 29.313L19 33l-3.812 8.313l-.25 11L15 56h42l-.312-15.437L53 33l-6.062-4L46 27.625l-20.312.938z"/><path fill="#d22f27" d="M26.75 25L26 29v1h20.125l-.844-5.469l-2.093-.281l-1.5 2.188l-11.625.062l-1.25-2.062l-2.157.156z"/><path fill="#d22f27" d="M39.344 29.844S52.25 34.25 52 55.75c6.375 2.063 5.5-6.5 5.5-6.5l-2.583-13.417l-5.75-6.25l-9.823.26z"/><circle cx="30" cy="35" r="2"/><circle cx="30" cy="41" r="2"/><circle cx="30" cy="47" r="2"/><circle cx="36" cy="35" r="2"/><circle cx="36" cy="41" r="2"/><circle cx="36" cy="47" r="2"/><circle cx="42" cy="35" r="2"/><circle cx="42" cy="41" r="2"/><circle cx="42" cy="47" r="2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M26 24v5s-3.5.5-7 4s-4 10-4 10v13h42V43.125S56.5 36.5 53 33s-7-4-7-4v-5h-3l-1 3H29.813L29 24h-3z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M61.984 16.452c-8-6-44-6-52 0c0 0-4 4-2 10h13v-5c6-4 24-4 30 0v5h13c2-6-2-10-2-10z"/>`),
		g.Group(children),
	)
}