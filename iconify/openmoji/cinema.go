package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cinema(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<rect width="50" height="50" x="11" y="10.958" fill="#92d3f5" rx="1.921"/><path fill="#fff" d="m43.96 46.28l7.63 6.22a1.09 1.09 0 0 0 1.778-.845V33.477a1.09 1.09 0 0 0-1.778-.844l-7.84 6.39"/><path fill="#fff" d="M41.773 52.486H20.82a2.188 2.188 0 0 1-2.188-2.188V34.833a2.188 2.188 0 0 1 2.188-2.188h20.953a2.188 2.188 0 0 1 2.187 2.188v15.465a2.188 2.188 0 0 1-2.187 2.188Z"/><circle cx="23.647" cy="24.253" r="5" fill="#fff"/><circle cx="38.749" cy="24.253" r="5" fill="#fff"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="m43.96 46.28l7.63 6.22a1.09 1.09 0 0 0 1.778-.845V33.477a1.09 1.09 0 0 0-1.778-.844l-7.63 6.18"/><circle cx="23.647" cy="24.253" r="5" stroke-linejoin="round"/><circle cx="38.749" cy="24.253" r="5" stroke-linejoin="round"/><path stroke-linejoin="round" d="M41.773 52.486H20.82a2.188 2.188 0 0 1-2.188-2.188V34.833a2.188 2.188 0 0 1 2.188-2.188h20.953a2.188 2.188 0 0 1 2.187 2.188v15.465a2.188 2.188 0 0 1-2.187 2.188Z"/><rect width="50" height="50" x="11" y="10.958" stroke-miterlimit="10" rx="1.921"/></g>`),
		g.Group(children),
	)
}