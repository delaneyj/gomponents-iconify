package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepForwardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 5h3v14H7V5m5 0l11 7l-11 7V5M2 5h3v14H2V5Z"/>`),
		g.Group(children),
	)
}