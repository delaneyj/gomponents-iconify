package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColumnsDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 64v120H40V64Z" opacity=".2"/><path d="M120 64a8 8 0 0 1-8 8H40a8 8 0 0 1 0-16h72a8 8 0 0 1 8 8Zm-8 32H40a8 8 0 0 0 0 16h72a8 8 0 0 0 0-16Zm0 40H40a8 8 0 0 0 0 16h72a8 8 0 0 0 0-16Zm0 40H40a8 8 0 0 0 0 16h72a8 8 0 0 0 0-16Zm32-104h72a8 8 0 0 0 0-16h-72a8 8 0 0 0 0 16Zm72 24h-72a8 8 0 0 0 0 16h72a8 8 0 0 0 0-16Zm0 40h-72a8 8 0 0 0 0 16h72a8 8 0 0 0 0-16Zm0 40h-72a8 8 0 0 0 0 16h72a8 8 0 0 0 0-16Z"/></g>`),
		g.Group(children),
	)
}