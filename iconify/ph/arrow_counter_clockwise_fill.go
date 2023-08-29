package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCounterClockwiseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 128a96 96 0 0 1-94.71 96H128a95.38 95.38 0 0 1-65.9-26.2a8 8 0 0 1 11-11.63a80 80 0 1 0-1.67-114.78a3.07 3.07 0 0 1-.26.25l-10.54 9.65l17 17A8 8 0 0 1 72 112H24a8 8 0 0 1-8-8V56a8 8 0 0 1 13.66-5.7L49.31 70l10.94-10A96 96 0 0 1 224 128Z"/>`),
		g.Group(children),
	)
}