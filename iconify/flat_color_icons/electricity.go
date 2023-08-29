package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Electricity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#00BCD4" d="M33.7 5L22 17l15 5l-15.7 14.7l5.1 2.8L12 43l2.7-14.8l2.9 5.1L27 24l-15-5L25 5h8.7z"/>`),
		g.Group(children),
	)
}