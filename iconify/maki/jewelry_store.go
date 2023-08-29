package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JewelryStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M12 8.5a4.5 4.5 0 1 1-9 0a4.498 4.498 0 0 1 2.71-4.125l.176.137l.774.601A3.498 3.498 0 0 0 4 8.5C4 10.43 5.57 12 7.5 12S11 10.43 11 8.5a3.498 3.498 0 0 0-2.66-3.387l.95-.738A4.498 4.498 0 0 1 12 8.5zm-4.5-4L10 2.555L9 1H6L5 2.555l1.5 1.167l1 .778z"/>`),
		g.Group(children),
	)
}