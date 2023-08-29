package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ratings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#42A5F5" d="M36 44H8V8h20l8 8z"/><path fill="#90CAF9" d="M40 40H12V4h20l8 8z"/><path fill="#E1F5FE" d="M38.5 13H31V5.5z"/><path fill="#1976D2" d="M34 20h-7l2.4 2.4l-2.4 2.5l-4-4l-6.1 6l2.2 2.2l3.9-4l4 4l4.6-4.5L34 27z"/>`),
		g.Group(children),
	)
}