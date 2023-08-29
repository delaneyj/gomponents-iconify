package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 11V2H2v9H0v2h16v-2h-2zm-4 1H6v-1h4v1zm3-2H3V3h10v7z"/>`),
		g.Group(children),
	)
}