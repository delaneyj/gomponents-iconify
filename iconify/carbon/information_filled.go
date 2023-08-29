package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InformationFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="M16 8a1.5 1.5 0 1 1-1.5 1.5A1.5 1.5 0 0 1 16 8Zm4 13.875h-2.875v-8H13v2.25h1.875v5.75H12v2.25h8Z"/><path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14 14 0 0 0 16 2Zm0 6a1.5 1.5 0 1 1-1.5 1.5A1.5 1.5 0 0 1 16 8Zm4 16.125h-8v-2.25h2.875v-5.75H13v-2.25h4.125v8H20Z"/>`),
		g.Group(children),
	)
}