package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBottomFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 216a8 8 0 0 1-8 8H40a8 8 0 0 1 0-16h176a8 8 0 0 1 8 8Zm-72-24h40a16 16 0 0 0 16-16V80a16 16 0 0 0-16-16h-40a16 16 0 0 0-16 16v96a16 16 0 0 0 16 16Zm-88 0h40a16 16 0 0 0 16-16V40a16 16 0 0 0-16-16H64a16 16 0 0 0-16 16v136a16 16 0 0 0 16 16Z"/>`),
		g.Group(children),
	)
}