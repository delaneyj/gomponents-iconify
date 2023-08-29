package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CouponFiveLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.005 14v7a1 1 0 0 1-1 1h-16a1 1 0 0 1-1-1v-7a2 2 0 1 0 0-4V3a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v7a2 2 0 1 0 0 4Zm-2 1.465a3.998 3.998 0 0 1-2-3.465c0-1.48.804-2.773 2-3.465V4h-14v4.535a4 4 0 0 1 0 6.93V20h14v-4.535ZM9.005 6h6v2h-6V6Zm0 10h6v2h-6v-2Z"/>`),
		g.Group(children),
	)
}