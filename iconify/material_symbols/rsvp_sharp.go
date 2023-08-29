package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RsvpSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.25 15L12.5 9H14l1 3.425L16 9h1.5l-1.75 6h-1.5ZM1 15V9h5v4h-.85L6 15H4.5l-.85-2H2.5v2H1Zm1.5-3.5h2v-1h-2v1ZM18 15V9h5v4h-3.5v2H18Zm1.5-3.5h2v-1h-2v1ZM7 15v-1.5h3v-.75H7V9h4.5v1.5h-3v.75h3V15H7Z"/>`),
		g.Group(children),
	)
}