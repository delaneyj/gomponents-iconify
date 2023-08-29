package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelRightCursorTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 3a3 3 0 0 1 3 3v7a2.99 2.99 0 0 1-.375 1.454L12.56 9.435a1.5 1.5 0 0 0-1.06-.434V4H5a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h5.002v1H5a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3h10Zm-3.144 7.146a.5.5 0 0 0-.852.354l-.004 8a.5.5 0 0 0 .9.301l1.994-2.647l3.497.776a.5.5 0 0 0 .46-.843l-5.995-5.94Z"/>`),
		g.Group(children),
	)
}