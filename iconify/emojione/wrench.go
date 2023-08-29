package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wrench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#94989b" d="m47.4 47.4l7.5-2l8.8 8.8c1-4.2-.1-8.9-3.4-12.2c-3.3-3.3-7.9-4.4-12.2-3.4L25.4 16c1-4.2-.1-8.9-3.4-12.2C18.7.5 14.1-.6 9.8.4l8.8 8.8l-2 7.5l-7.5 2L.4 9.8c-1 4.2.1 8.9 3.4 12.2c3.3 3.3 7.9 4.4 12.2 3.4L38.6 48c-1 4.2.1 8.9 3.4 12.2c3.3 3.3 7.9 4.4 12.2 3.4l-8.8-8.8l2-7.4"/>`),
		g.Group(children),
	)
}