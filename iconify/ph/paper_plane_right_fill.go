package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 127.89a16 16 0 0 1-8.18 14L55.91 237.9A16.14 16.14 0 0 1 48 240a16 16 0 0 1-15.05-21.34l27.35-79.95a4 4 0 0 1 3.79-2.71H136a8 8 0 0 0 8-8.53a8.19 8.19 0 0 0-8.26-7.47H64.16a4 4 0 0 1-3.79-2.7l-27.44-80a16 16 0 0 1 22.92-19.23l168 95.89a16 16 0 0 1 8.15 13.93Z"/>`),
		g.Group(children),
	)
}