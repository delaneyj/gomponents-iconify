package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ladder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 1v2h-4V1H8v21h2v-2h4v2h2V1h-2m0 4v3h-4V5h4m0 5v3h-4v-3h4m-4 8v-3h4v3h-4Z"/>`),
		g.Group(children),
	)
}