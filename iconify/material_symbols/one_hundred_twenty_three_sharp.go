package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneHundredTwentyThreeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 15v-4.5H4V9h3v6H5.5ZM9 15v-3.5h3v-1H9V9h4.5v3.5h-3v1h3V15H9Zm6 0v-1.5h3v-1h-2v-1h2v-1h-3V9h4.5v6H15Z"/>`),
		g.Group(children),
	)
}