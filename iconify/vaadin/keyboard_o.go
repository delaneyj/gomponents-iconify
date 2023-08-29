package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 5v7H1V5h14zm1-1H0v9h16V4z"/><path fill="currentColor" d="M4 10h8v1H4v-1zm-2 0h1v1H2v-1zm11 0h1v1h-1v-1zm-2-2h1v1h-1V8zM9 8h1v1H9V8zM7 8h1v1H7V8zM5 8h1v1H5V8zM3 8h1v1H3V8zm7-2h1v1h-1V6zm2 0v1h1v2h1V6zM8 6h1v1H8V6zM6 6h1v1H6V6zM4 6h1v1H4V6zM2 6h1v1H2V6z"/>`),
		g.Group(children),
	)
}