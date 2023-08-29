package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScrollHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2v2H2V2h20zm0 18v2H2v-2h20zm-6-5v-2H8v2H6v-2H4v-2h2V9h2v2h8V9h2v2h2v2h-2v2h-2zm0 0v2h-2v-2h2zm0-6h-2V7h2v2zM8 9V7h2v2H8zm0 6h2v2H8v-2z"/>`),
		g.Group(children),
	)
}