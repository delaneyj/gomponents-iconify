package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashAuto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#FFC107" d="M33 22h-9.4L30 5H19l-6 21h8.6L17 45z"/><path fill="#F44336" d="M40.8 14.5h-4.3l-.9 2.5H33l4.5-12h2.3l4.5 12h-2.6l-.9-2.5zm-3.7-2h3L38.6 8l-1.5 4.5z"/>`),
		g.Group(children),
	)
}