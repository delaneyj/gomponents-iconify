package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTopFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 40a8 8 0 0 1-8 8H40a8 8 0 0 1 0-16h176a8 8 0 0 1 8 8Zm-32 24h-40a16 16 0 0 0-16 16v96a16 16 0 0 0 16 16h40a16 16 0 0 0 16-16V80a16 16 0 0 0-16-16Zm-88 0H64a16 16 0 0 0-16 16v136a16 16 0 0 0 16 16h40a16 16 0 0 0 16-16V80a16 16 0 0 0-16-16Z"/>`),
		g.Group(children),
	)
}