package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 6.594l-.719.687l-12.5 12.5L4.22 21.22L16 9.437L27.781 21.22l1.438-1.438l-12.5-12.5z"/>`),
		g.Group(children),
	)
}