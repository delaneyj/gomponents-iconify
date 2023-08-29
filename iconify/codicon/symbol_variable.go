package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymbolVariable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5h2V4H1.5l-.5.5v8l.5.5H4v-1H2V5zm12.5-1H12v1h2v7h-2v1h2.5l.5-.5v-8l-.5-.5zm-2.74 2.57L12 7v2.51l-.3.45l-4.5 2h-.46l-2.5-1.5l-.24-.43v-2.5l.3-.46l4.5-2h.46l2.5 1.5zM5 9.71l1.5.9V9.28L5 8.38v1.33zm.58-2.15l1.45.87l3.39-1.5l-1.45-.87l-3.39 1.5zm1.95 3.17l3.5-1.56v-1.4l-3.5 1.55v1.41z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}