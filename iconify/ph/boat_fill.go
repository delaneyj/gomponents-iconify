package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoatFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M221.06 110.63L208 106.27V56a16 16 0 0 0-16-16h-56V24a8 8 0 0 0-16 0v16H64a16 16 0 0 0-16 16v50.27l-13.06 4.36A16 16 0 0 0 24 125.82v34.28a7.73 7.73 0 0 0 .31 2.19c15.72 55.09 86.12 74.06 100.07 77.3a16.19 16.19 0 0 0 7.24 0c13.95-3.24 84.35-22.21 100.07-77.3a7.73 7.73 0 0 0 .31-2.19v-34.28a16 16 0 0 0-10.94-15.19ZM136 168a8 8 0 0 1-16 0v-56a8 8 0 0 1 16 0Zm56-67.07l-61.47-20.52a8 8 0 0 0-5.06 0L64 100.93V56h128Z"/>`),
		g.Group(children),
	)
}