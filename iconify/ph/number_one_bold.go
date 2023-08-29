package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberOneBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M148 48v160a12 12 0 0 1-24 0V69.19l-21.83 13.1a12 12 0 0 1-12.34-20.58l40-24A12 12 0 0 1 148 48Z"/>`),
		g.Group(children),
	)
}