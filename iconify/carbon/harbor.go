package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Harbor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 17a11.01 11.01 0 0 1-10 10.95V14h6v-2h-6V9.858a4 4 0 1 0-2 0V12H9v2h6v13.95A11.01 11.01 0 0 1 5 17H3a13 13 0 0 0 26 0ZM14 6a2 2 0 1 1 2 2a2.002 2.002 0 0 1-2-2Z"/>`),
		g.Group(children),
	)
}