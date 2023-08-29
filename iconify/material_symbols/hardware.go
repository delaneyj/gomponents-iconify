package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hardware(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 11V8H4q0-2.075 1.463-3.538T9 3h6v3l3-3h2v8h-2l-3-3v3H9Zm1 10q-.425 0-.713-.288T9 20v-7h6v7q0 .425-.288.713T14 21h-4Z"/>`),
		g.Group(children),
	)
}