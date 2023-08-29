package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VariableSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.034 2.234a.8.8 0 0 1 1.132 0L10.4 6.47l4.234-4.235a.8.8 0 0 1 1.132 1.132l-4.8 4.8a.8.8 0 0 1-1.132 0L5.6 3.93L1.366 8.166A.8.8 0 0 1 .234 7.034l4.8-4.8ZM1.6 12.667v-1.334a.8.8 0 0 0-1.6 0v2.934h16v-2.934a.8.8 0 0 0-1.6 0v1.334h-3.2v-1.334a.8.8 0 0 0-1.6 0v1.334H6.4v-1.334a.8.8 0 0 0-1.6 0v1.334H1.6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}