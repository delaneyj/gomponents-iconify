package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewWeekSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h5.325v16H2Zm7.35 0V4h5.325v16H9.35Zm7.325 0V4H22v16h-5.325Z"/>`),
		g.Group(children),
	)
}