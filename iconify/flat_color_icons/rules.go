package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rules(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#42A5F5" d="M39 45H9s-3-.1-3-8h36c0 7.9-3 8-3 8z"/><path fill="#90CAF9" d="M8 3h32v34H8z"/><path fill="#1976D2" d="M18 15h16v2H18zm0 4h16v2H18zm0 4h16v2H18zm0 4h16v2H18zm0 4h16v2H18zm-4-16h2v2h-2zm0 4h2v2h-2zm0 4h2v2h-2zm0 4h2v2h-2zm0 4h2v2h-2z"/>`),
		g.Group(children),
	)
}