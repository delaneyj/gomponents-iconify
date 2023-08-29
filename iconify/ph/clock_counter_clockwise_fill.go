package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockCounterClockwiseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 128a96 96 0 0 1-161.89 69.82a8 8 0 1 1 11-11.64a80 80 0 1 0-1.68-114.75A439.19 439.19 0 0 0 61.35 82l16.31 16.34A8 8 0 0 1 72 112H32a8 8 0 0 1-8-8V64a8 8 0 0 1 13.66-5.66L50 70.7c3.22-3.49 6.54-7 10.06-10.55A96 96 0 0 1 224 128Zm-96-56a8 8 0 0 0-8 8v48a8 8 0 0 0 3.88 6.86l40 24a8 8 0 1 0 8.24-13.72L136 123.47V80a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}