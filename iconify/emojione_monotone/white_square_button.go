package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteSquareButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M2 2v60h60V2H2zm58 58H4V4h56v56z"/><path fill="currentColor" d="M10 10h44v44H10z"/>`),
		g.Group(children),
	)
}