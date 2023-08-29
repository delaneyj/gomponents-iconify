package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortPrecedence(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 2L6 3v3h1V3h7v2.453l.207-.16l.793.793V3l-1-1H7zm1 2h2v2H8V4zM5 9H3v2h2V9zM2 7L1 8v5l1 1h7l1-1V8L9 7H2zm0 6V8h7v5H2zm6-3H6v2h2v-2zm5-6h-1v3.864l-1.182-1.182l-.707.707l2.035 2.036h.708l2.035-2.036l-.707-.707L13 7.864V4z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}