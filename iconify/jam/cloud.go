package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 14a5 5 0 1 1 1.561-9.751A7.002 7.002 0 0 1 20 7a7 7 0 0 1-7 7H5zm0-2h8a5 5 0 1 0-4.6-6.965l-.72 1.686l-1.742-.572A3 3 0 1 0 5 12z"/>`),
		g.Group(children),
	)
}