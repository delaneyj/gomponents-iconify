package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodiumGold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 7.09l2.45 1.49l-.65-2.81L16 3.89l-2.89-.25L12 1l-1.13 2.64L8 3.89l2.18 1.88l-.68 2.81L12 7.09M15 23H9V10h6v13M1 17v6h6v-6H1m4 4H3v-2h2v2m12-8v10h6V13h-6m4 8h-2v-6h2v6Z"/>`),
		g.Group(children),
	)
}