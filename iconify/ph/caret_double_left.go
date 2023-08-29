package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDoubleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M205.66 202.34a8 8 0 0 1-11.32 11.32l-80-80a8 8 0 0 1 0-11.32l80-80a8 8 0 0 1 11.32 11.32L131.31 128ZM51.31 128l74.35-74.34a8 8 0 0 0-11.32-11.32l-80 80a8 8 0 0 0 0 11.32l80 80a8 8 0 0 0 11.32-11.32Z"/>`),
		g.Group(children),
	)
}