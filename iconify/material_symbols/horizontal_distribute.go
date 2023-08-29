package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalDistribute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V2h2v20H2Zm8.5-5V7h3v10h-3Zm9.5 5V2h2v20h-2Z"/>`),
		g.Group(children),
	)
}