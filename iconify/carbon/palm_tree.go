package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PalmTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18.57 28C17.897 26.725 16 22.226 16 12v-.131l1.117.744A6.461 6.461 0 0 1 20 18h2a8.457 8.457 0 0 0-3.773-7.05L16.803 10h1.53a7.04 7.04 0 0 1 4.2 1.4l1.867 1.4l1.2-1.6l-1.867-1.4a9.06 9.06 0 0 0-5.4-1.8h-1.216A7.03 7.03 0 0 1 22 6h2V4h-2a9.035 9.035 0 0 0-7 3.364A9.035 9.035 0 0 0 8 4H6v2h2a7.03 7.03 0 0 1 4.883 2h-1.216a9.06 9.06 0 0 0-5.4 1.8L4.4 11.2l1.2 1.6l1.867-1.4a7.04 7.04 0 0 1 4.2-1.4h1.53l-1.423.95A8.457 8.457 0 0 0 8 18h2a6.461 6.461 0 0 1 2.883-5.387L14 11.87V12c0 8.944 1.412 13.765 2.361 16H2v2h28v-2Z"/>`),
		g.Group(children),
	)
}