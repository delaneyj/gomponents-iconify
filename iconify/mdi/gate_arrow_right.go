package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GateArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 6v5h-2V7h-2v4H9V9H7v12h2v-2h2v2h1.09c-.06-.33-.09-.66-.09-1c0-1.18.35-2.33 1-3.31V13h2v1.81c.62-.36 1.3-.6 2-.72V13h2v1.09c.7.12 1.38.36 2 .72V13h1v-2h-1V6h-2v5h-2V6h-2m-6 7h2v4H9v-4m10 4v2h-4v2h4v2l3-3l-3-3Z"/>`),
		g.Group(children),
	)
}