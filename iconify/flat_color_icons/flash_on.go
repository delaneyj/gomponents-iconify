package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#FFC107" d="M33 22h-9.4L30 5H19l-6 21h8.6L17 45z"/>`),
		g.Group(children),
	)
}