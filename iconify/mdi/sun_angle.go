package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunAngle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.8 5.2C13 5 12.2 5 11.4 5l3.2-2.7l1.4 4c-.7-.5-1.4-.8-2.2-1.1M7 7.1c.6-.6 1.3-1.1 2-1.4l-4.1-.8l.7 4.1c.4-.7.8-1.4 1.4-1.9m-1.8 6.7C5 13 5 12.2 5 11.4l-2.7 3.2l4 1.4c-.5-.6-.9-1.4-1.1-2.2M22 19v2H3l5.4-5.5c-1.9-2-1.9-5.1 0-7.1c1.9-1.9 5.1-1.9 7 0l3-3l1.4 1.4L7.7 19H22Z"/>`),
		g.Group(children),
	)
}