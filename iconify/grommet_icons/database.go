package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Database(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 2h22v7H1V2Zm3 10h1v1H4v-1Zm0-7h1v1H4V5Zm0 14h1v1H4v-1Zm-3-3h22v7H1v-7Zm0-7h22v7H1V9Z"/>`),
		g.Group(children),
	)
}