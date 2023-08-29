package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymbolParameter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 6h-1v-.5a.5.5 0 0 0-.5-.5H8.479v5.5a.5.5 0 0 0 .5.5h.5v1h-3v-1h.5a.5.5 0 0 0 .5-.5V5H6.5a.5.5 0 0 0-.5.5V6H5V4h6v2zm2.914 2.048l-1.462-1.462l.707-.707l1.816 1.816v.707l-1.768 1.767l-.707-.707l1.414-1.414zM3.548 9.462L2.086 8L3.5 6.586l-.707-.707l-1.768 1.767v.708l1.816 1.815l.707-.707z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}