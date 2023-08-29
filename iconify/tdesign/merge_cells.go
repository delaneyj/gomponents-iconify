package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MergeCells(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h7v-4h2v4h7V4h-7v4h-2V4H4Zm13.182 6.232L15.414 12l1.768 1.768l-1.414 1.414L12.586 12l3.182-3.182l1.414 1.414ZM8.33 8.818L11.512 12L8.33 15.182l-1.414-1.414L8.683 12l-1.767-1.768L8.33 8.818Z"/>`),
		g.Group(children),
	)
}