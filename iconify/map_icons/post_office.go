package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PostOffice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M2.941 8c-2.941 0-1.47.779 0 1.974L25 26.763l22.059-16.737C48.531 8.831 50 8 47.059 8H2.941zM0 11.946v24.728C0 38.129 1.488 40 2.665 40h44.67C48.513 40 50 38.129 50 36.674V11.946L25 31.021L0 11.946z"/>`),
		g.Group(children),
	)
}