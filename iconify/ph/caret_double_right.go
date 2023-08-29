package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDoubleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m141.66 133.66l-80 80a8 8 0 0 1-11.32-11.32L124.69 128L50.34 53.66a8 8 0 0 1 11.32-11.32l80 80a8 8 0 0 1 0 11.32Zm80-11.32l-80-80a8 8 0 0 0-11.32 11.32L204.69 128l-74.35 74.34a8 8 0 0 0 11.32 11.32l80-80a8 8 0 0 0 0-11.32Z"/>`),
		g.Group(children),
	)
}