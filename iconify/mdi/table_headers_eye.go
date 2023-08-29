package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableHeadersEye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 6v13a2 2 0 0 0 2 2h6V11h12V6a2 2 0 0 0-2-2H3a2 2 0 0 0-2 2m2 0h4v3H3m4 10H3v-3h4m0-2H3v-3h4m2-2V6h4v3m6 0h-4V6h4m-2 10a1 1 0 1 1-1 1a1 1 0 0 1 1-1m0-3a6.45 6.45 0 0 1 6 4a6.5 6.5 0 0 1-12 0a6.45 6.45 0 0 1 6-4m0 1.5a2.5 2.5 0 1 0 2.5 2.5a2.5 2.5 0 0 0-2.5-2.5"/>`),
		g.Group(children),
	)
}