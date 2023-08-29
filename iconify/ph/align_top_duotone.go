package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTopDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M200 80v96a8 8 0 0 1-8 8h-40a8 8 0 0 1-8-8V80a8 8 0 0 1 8-8h40a8 8 0 0 1 8 8Zm-96-8H64a8 8 0 0 0-8 8v136a8 8 0 0 0 8 8h40a8 8 0 0 0 8-8V80a8 8 0 0 0-8-8Z" opacity=".2"/><path d="M224 40a8 8 0 0 1-8 8H40a8 8 0 0 1 0-16h176a8 8 0 0 1 8 8Zm-16 40v96a16 16 0 0 1-16 16h-40a16 16 0 0 1-16-16V80a16 16 0 0 1 16-16h40a16 16 0 0 1 16 16Zm-16 0h-40v96h40Zm-72 0v136a16 16 0 0 1-16 16H64a16 16 0 0 1-16-16V80a16 16 0 0 1 16-16h40a16 16 0 0 1 16 16Zm-16 0H64v136h40Z"/></g>`),
		g.Group(children),
	)
}