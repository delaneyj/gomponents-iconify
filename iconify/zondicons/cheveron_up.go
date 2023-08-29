package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheveronUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.707 7.05L10 6.343L4.343 12l1.414 1.414L10 9.172l4.243 4.242L15.657 12z"/>`),
		g.Group(children),
	)
}