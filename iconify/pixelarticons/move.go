package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Move(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 0h-2v2H9v2H7v2h2V4h2v7H4V9h2V7H4v2H2v2H0v2h2v2h2v2h2v-2H4v-2h7v7H9v-2H7v2h2v2h2v2h2v-2h2v-2h2v-2h-2v2h-2v-7h7v2h-2v2h2v-2h2v-2h2v-2h-2V9h-2V7h-2v2h2v2h-7V4h2v2h2V4h-2V2h-2V0z"/>`),
		g.Group(children),
	)
}