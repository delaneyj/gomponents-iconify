package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignJustifyThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M36 64a4 4 0 0 1 4-4h176a4 4 0 0 1 0 8H40a4 4 0 0 1-4-4Zm180 36H40a4 4 0 0 0 0 8h176a4 4 0 0 0 0-8Zm0 40H40a4 4 0 0 0 0 8h176a4 4 0 0 0 0-8Zm0 40H40a4 4 0 0 0 0 8h176a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}