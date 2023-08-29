package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardwareOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 21q-.425 0-.713-.288T9 20V8H4q0-2.075 1.463-3.538T9 3h6v3l3-3h2v8h-2l-3-3v12q0 .425-.288.713T14 21h-4Zm1-2h2v-6h-2v6Zm0-8h2V5H9q-.65 0-1.225.263t-1 .737H11v5Zm1 1Z"/>`),
		g.Group(children),
	)
}