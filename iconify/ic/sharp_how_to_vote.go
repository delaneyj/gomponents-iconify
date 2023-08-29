package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpHowToVote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 13h-.68l-2 2h1.91L19 17H5l1.78-2h2.05l-2-2H6l-3 3v6h18v-6zm1.81-5.04L13.45 1.6L5.68 9.36l6.36 6.36l7.77-7.76zm-6.35-3.55L17 7.95l-4.95 4.95l-3.54-3.54l4.95-4.95z"/>`),
		g.Group(children),
	)
}