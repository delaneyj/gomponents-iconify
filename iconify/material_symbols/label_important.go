package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelImportant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3 19l4.5-7L3 5h12q.5 0 .938.225t.712.625L21 12l-4.35 6.15q-.275.4-.713.625T15 19H3Z"/>`),
		g.Group(children),
	)
}