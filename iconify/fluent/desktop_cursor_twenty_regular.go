package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopCursorTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 2a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h3v2H5.5a.5.5 0 0 0 0 1H11l.001-1H8v-2h3.002l.001-1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v8.843l.872.864A2.02 2.02 0 0 0 18 13V4a2 2 0 0 0-2-2H4Zm8.857 8.146a.5.5 0 0 0-.852.354l-.005 8a.5.5 0 0 0 .9.301l1.995-2.647l3.496.776a.5.5 0 0 0 .46-.843l-5.994-5.94Z"/>`),
		g.Group(children),
	)
}