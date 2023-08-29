package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.2 1.2h-2.4v9.6H1.2v2.4h9.6v9.6h2.4v-9.6h9.6v-2.4h-9.6z"/>`),
		g.Group(children),
	)
}