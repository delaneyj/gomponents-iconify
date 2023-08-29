package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 7h4v2H7a3 3 0 0 0-3 3a3 3 0 0 0 3 3h4v2H7a5 5 0 0 1-5-5a5 5 0 0 1 5-5m10 0a5 5 0 0 1 5 5h-2a3 3 0 0 0-3-3h-4V7h4m-9 4h8v2H8v-2m9 1h2v3h3v2h-3v3h-2v-3h-3v-2h3v-3Z"/>`),
		g.Group(children),
	)
}