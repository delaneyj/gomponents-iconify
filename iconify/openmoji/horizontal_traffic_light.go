package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalTrafficLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g stroke-miterlimit="10"><path fill="#d0cfce" d="M54.09 46.91h-36c-3.85 0-7-3.15-7-7v-8c0-3.85 3.15-7 7-7h36c3.85 0 7 3.15 7 7v8c0 3.85-3.15 7-7 7z"/><circle cx="21.618" cy="35.91" r="5" fill="#b1cc33"/><circle cx="35.618" cy="35.91" r="5" fill="#f4aa41"/><circle cx="50.618" cy="35.91" r="5" fill="#ea5a47"/></g><g fill="none" stroke="#000" stroke-miterlimit="10"><path d="M54.09 46.91h-36c-3.85 0-7-3.15-7-7v-8c0-3.85 3.15-7 7-7h36c3.85 0 7 3.15 7 7v8c0 3.85-3.15 7-7 7z"/><circle cx="21.618" cy="35.91" r="5"/><circle cx="35.618" cy="35.91" r="5"/><circle cx="50.618" cy="35.91" r="5"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M53.91 46.91h-36c-3.85 0-7-3.15-7-7v-8c0-3.85 3.15-7 7-7h36c3.85 0 7 3.15 7 7v8c0 3.85-3.15 7-7 7z"/><circle cx="21.438" cy="35.91" r="5" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><circle cx="35.438" cy="35.91" r="5" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><circle cx="50.438" cy="35.91" r="5" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/></g>`),
		g.Group(children),
	)
}