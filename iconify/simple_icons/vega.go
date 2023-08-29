package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vega(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.39 8.693H24l-3.123-6.68ZM6.68 1.987H0L7.098 22.03h.008l2.86-10.73zm14.197-.016l-7.098 20.042h-6.68L14.195 1.97"/>`),
		g.Group(children),
	)
}