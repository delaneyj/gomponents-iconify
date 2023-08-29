package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighVoltage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="m45 2l-9.396 2.48L29.298 2L19 36.354h10.865L20.352 62L43.65 27.648H29.626z"/>`),
		g.Group(children),
	)
}