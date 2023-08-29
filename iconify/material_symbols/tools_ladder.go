package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsLadder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.95 21q-.475 0-.75-.388t-.15-.837L9.5 3.7q.1-.325.35-.512T10.425 3q.5 0 .775.388t.15.837L10.85 6h5.625l.625-2.3q.1-.325.363-.512T18.05 3q.475 0 .75.388t.15.837L14.5 20.3q-.1.325-.35.513t-.575.187q-.5 0-.775-.388t-.15-.837l.5-1.775H7.525L6.9 20.3q-.1.325-.363.513T5.95 21Zm3.525-10h5.6l.825-3h-5.6l-.825 3ZM8.1 16h5.6l.825-3h-5.6L8.1 16Z"/>`),
		g.Group(children),
	)
}