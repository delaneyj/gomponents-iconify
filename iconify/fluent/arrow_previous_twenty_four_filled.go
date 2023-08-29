package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowPreviousTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 3a1 1 0 0 0-.993.883L5 4v16a1 1 0 0 0 1.993.117L7 20V4a1 1 0 0 0-1-1Zm12.707.293a1 1 0 0 0-1.32-.083l-.094.083l-8 8a1 1 0 0 0-.083 1.32l.083.094l8 8a1 1 0 0 0 1.497-1.32l-.083-.094L11.414 12l7.293-7.293a1 1 0 0 0 0-1.414Z"/>`),
		g.Group(children),
	)
}