package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymbolField(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m14.45 4.5l-5-2.5h-.9l-7 3.5l-.55.89v4.5l.55.9l5 2.5h.9l7-3.5l.55-.9v-4.5l-.55-.89zm-8 8.64l-4.5-2.25V7.17l4.5 2v3.97zm.5-4.8L2.29 6.23l6.66-3.34l4.67 2.34l-6.67 3.11zm7 1.55l-6.5 3.25V9.21l6.5-3v3.68z"/>`),
		g.Group(children),
	)
}