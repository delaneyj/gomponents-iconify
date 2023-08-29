package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileExcel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V9.6l-.3-.3l-6-6l-.3-.3H6zm2 2h10v6h6v16H8V5zm12 1.4L22.6 9H20V6.4zM11 13l3.8 5.5L11 24h2.4l2.6-3.8l2.6 3.8H21l-3.8-5.5L21 13h-2.4L16 16.8L13.4 13H11z"/>`),
		g.Group(children),
	)
}