package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InputPower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M214 85q15 0 28.5 13.5T256 128v117l-75 75v64H75v-64L0 245V128q0-16 13.5-29.5T42 85h1V0h42v85h86V0h42z"/>`),
		g.Group(children),
	)
}