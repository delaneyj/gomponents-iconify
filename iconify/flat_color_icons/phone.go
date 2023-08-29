package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#009688" d="M39.1 7h-3.7C22.2 7.2 7.1 24.1 7 35.4v3.7c0 1 .8 1.9 1.9 1.9l7.5-.1c1 0 1.9-.9 1.9-1.9l.2-8.2l-4.7-4c0-2.6 10.5-13.1 13.2-13.2l4.3 4.7l7.9-.2c1 0 1.9-.9 1.9-1.9L41 8.9c0-1.1-.8-1.9-1.9-1.9z"/>`),
		g.Group(children),
	)
}