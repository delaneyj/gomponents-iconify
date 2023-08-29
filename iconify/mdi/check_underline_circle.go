package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckUnderlineCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2A10 10 0 0 0 2 12a10 10 0 0 0 10 10a10 10 0 0 0 10-10A10 10 0 0 0 12 2m5 16H7v-2h10v2m-6.7-4L7 10.7l1.4-1.4l1.9 1.9l5.3-5.3L17 7.3L10.3 14Z"/>`),
		g.Group(children),
	)
}