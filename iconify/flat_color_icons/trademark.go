package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trademark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21" fill="#9C27B0"/><path fill="#E1BEE7" d="M20.6 18.5h-4.2v14.2h-3.5V18.5H8.7v-2.9h11.9v2.9zm6.5-2.9L30.3 28l3.2-12.4H38v17.1h-3.5v-4.6l.3-7.1l-3.4 11.8H29L25.7 21l.3 7.1v4.6h-3.5V15.6h4.6z"/>`),
		g.Group(children),
	)
}