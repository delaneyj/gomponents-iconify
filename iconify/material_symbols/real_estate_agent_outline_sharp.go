package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RealEstateAgentOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 14V7.5L14 4L9 7.5V9H7V6.5l7-5l7 5V14h-2Zm-5 8.5l-7-1.95V22H1V11h7.95L17 14v2h5v4l-8 2.5ZM3 20h2v-7H3v7Zm10.95.4l5.95-1.85V18h-7.075L9.7 16.95l.6-1.9l2.925.95H15v-.65L8.6 13H7v5.5l6.95 1.9ZM14 4Zm.5 4h1V7h-1v1Zm-2 0h1V7h-1v1Zm2 2h1V9h-1v1Zm-2 0h1V9h-1v1Z"/>`),
		g.Group(children),
	)
}