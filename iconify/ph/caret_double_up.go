package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDoubleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M213.66 202.34a8 8 0 0 1-11.32 11.32L128 139.31l-74.34 74.35a8 8 0 0 1-11.32-11.32l80-80a8 8 0 0 1 11.32 0Zm-160-68.68L128 59.31l74.34 74.35a8 8 0 0 0 11.32-11.32l-80-80a8 8 0 0 0-11.32 0l-80 80a8 8 0 0 0 11.32 11.32Z"/>`),
		g.Group(children),
	)
}