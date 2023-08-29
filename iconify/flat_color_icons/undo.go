package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Undo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#00BCD4"><path d="M5 18L19 6.3v23.4z"/><path d="M28 14H16v8h12c2.8 0 5 2.2 5 5s-2.2 5-5 5h-3v8h3c7.2 0 13-5.8 13-13s-5.8-13-13-13z"/></g>`),
		g.Group(children),
	)
}