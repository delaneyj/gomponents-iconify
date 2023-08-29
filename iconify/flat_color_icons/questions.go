package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Questions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#42A5F5" d="M36 44H8V8h20l8 8z"/><path fill="#90CAF9" d="M40 40H12V4h20l8 8z"/><path fill="#E1F5FE" d="M38.5 13H31V5.5z"/><path fill="#1976D2" d="M24.5 28.3c0-4.7 3.6-4.4 3.6-7.2c0-.7-.2-2.1-2-2.1c-2 0-2.1 1.6-2.1 2h-2.7c0-.7.3-4.2 4.8-4.2c4.6 0 4.7 3.6 4.7 4.3c0 3.5-3.8 4-3.8 7.3h-2.5zm-.2 3.5c0-.2 0-1.5 1.5-1.5c1.4 0 1.5 1.3 1.5 1.5c0 .4-.2 1.4-1.5 1.4s-1.5-1-1.5-1.4z"/>`),
		g.Group(children),
	)
}