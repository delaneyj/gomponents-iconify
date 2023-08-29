package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerThreeSelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19v-3h6v-2.5H4v-3h6V8H4V5h6q1.25 0 2.125.875T13 8v1.9q0 .875-.613 1.488T10.9 12q.875 0 1.488.613T13 14.1V16q0 1.25-.875 2.125T10 19H4Zm11 0v-2h4v-1h-2.65q-.575 0-.963-.388T15 14.65v-2.3q0-.575.388-.963T16.35 11H21v2h-4v1h2.65q.575 0 .963.388t.387.962v2.3q0 .575-.388.963T19.65 19H15Z"/>`),
		g.Group(children),
	)
}