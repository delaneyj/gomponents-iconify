package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.425 21L3 19.575L7.6 15H5v-2h6v6H9v-2.6L4.425 21ZM13 11V5h2v2.6L19.575 3L21 4.425L16.4 9H19v2h-6Z"/>`),
		g.Group(children),
	)
}