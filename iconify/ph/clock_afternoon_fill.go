package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockAfternoonFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm56 112h-36.69l26.35 26.34a8 8 0 0 1-11.32 11.32l-40-40A8 8 0 0 1 128 120h56a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}