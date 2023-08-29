package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#FFC107" d="M33 22h-9.4L30 5H19l-6 21h8.6L17 45z"/><path fill="#37474F" d="M3.563 6.396L6.39 3.568l37.966 37.966l-2.828 2.828z"/>`),
		g.Group(children),
	)
}