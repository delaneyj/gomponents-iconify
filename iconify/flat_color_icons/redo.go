package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#00BCD4"><path d="M43 18L29 6.3v23.4z"/><path d="M20 14h12v8H20c-2.8 0-5 2.2-5 5s2.2 5 5 5h3v8h-3c-7.2 0-13-5.8-13-13s5.8-13 13-13z"/></g>`),
		g.Group(children),
	)
}