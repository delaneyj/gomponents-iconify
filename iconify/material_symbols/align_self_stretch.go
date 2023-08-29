package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignSelfStretch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4V2h20v2H2Zm0 18v-2h20v2H2Zm8.5-4.5V6h3v11.5h-3Z"/>`),
		g.Group(children),
	)
}