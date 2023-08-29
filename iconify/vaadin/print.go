package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Print(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 10v4h2v2h12v-2h2v-4H0zm13 5H3v-3h10v3zm-1-9V2L9.3 0H4v6H0v3h16V6h-4zM9 1l1.3 1H9V1zm2 6H5V1h3v2h3v4zm4 1h-1V7h1v1z"/>`),
		g.Group(children),
	)
}