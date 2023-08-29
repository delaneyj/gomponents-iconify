package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WbTwilight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.35 10.1l-1.4-1.45l2.15-2.1l1.4 1.4l-2.15 2.15ZM2 20v-2h20v2H2Zm9-13V4h2v3h-2Zm-5.35 3.05L3.55 7.9l1.4-1.4L7.1 8.65l-1.45 1.4ZM5 16q0-2.925 2.038-4.963T12 9q2.925 0 4.963 2.038T19 16H5Z"/>`),
		g.Group(children),
	)
}