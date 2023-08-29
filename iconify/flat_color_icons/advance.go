package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Advance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#1565C0"><path d="M46.1 24L33 35V13zM10 20h4v8h-4zm-6 0h4v8H4zm12 0h4v8h-4z"/><path d="M22 20h14v8H22z"/></g>`),
		g.Group(children),
	)
}