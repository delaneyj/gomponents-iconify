package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullCoverage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19V7h2v12h15v2H4Zm4-4q-.825 0-1.413-.588T6 15V3h17v12q0 .825-.588 1.413T21 17H8Zm2-5h4V7h-4v5Zm5 0h4v-2h-4v2Zm0-3h4V7h-4v2Z"/>`),
		g.Group(children),
	)
}