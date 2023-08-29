package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDoubleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M213.66 122.34a8 8 0 0 1 0 11.32l-80 80a8 8 0 0 1-11.32 0l-80-80a8 8 0 0 1 11.32-11.32L128 196.69l74.34-74.35a8 8 0 0 1 11.32 0Zm-91.32 11.32a8 8 0 0 0 11.32 0l80-80a8 8 0 0 0-11.32-11.32L128 116.69L53.66 42.34a8 8 0 0 0-11.32 11.32Z"/>`),
		g.Group(children),
	)
}