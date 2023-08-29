package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m0 15l6-5l-6-4.9zm9-4.9l6 4.9V5l-6 5.1zm5 2.8l-3.4-2.8l3.4-3v5.8zM7 5h1v1H7V5zm0-2h1v1H7V3zm0 4h1v1H7V7zm0 2h1v1H7V9zm0 2h1v1H7v-1zm0 2h1v1H7v-1zm0 2h1v1H7v-1z"/><path fill="currentColor" d="M7.5 1c1.3 0 2.6.7 3.6 1.9L10 4h3V1l-1.2 1.2C10.6.8 9.1 0 7.5 0C5.6 0 3.9 1 2.6 2.9l.8.6C4.5 1.9 5.9 1 7.5 1z"/>`),
		g.Group(children),
	)
}