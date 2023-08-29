package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepBackwardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 5h-3v14h3V5m-5 0L1 12l11 7V5m10 0h-3v14h3V5Z"/>`),
		g.Group(children),
	)
}