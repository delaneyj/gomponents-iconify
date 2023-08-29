package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21" fill="#4CAF50"/><g fill="#fff"><path d="M21 14h6v20h-6z"/><path d="M14 21h20v6H14z"/></g>`),
		g.Group(children),
	)
}