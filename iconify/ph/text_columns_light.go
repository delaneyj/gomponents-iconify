package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColumnsLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M118 64a6 6 0 0 1-6 6H40a6 6 0 0 1 0-12h72a6 6 0 0 1 6 6Zm-6 34H40a6 6 0 0 0 0 12h72a6 6 0 0 0 0-12Zm0 40H40a6 6 0 0 0 0 12h72a6 6 0 0 0 0-12Zm0 40H40a6 6 0 0 0 0 12h72a6 6 0 0 0 0-12Zm32-108h72a6 6 0 0 0 0-12h-72a6 6 0 0 0 0 12Zm72 28h-72a6 6 0 0 0 0 12h72a6 6 0 0 0 0-12Zm0 40h-72a6 6 0 0 0 0 12h72a6 6 0 0 0 0-12Zm0 40h-72a6 6 0 0 0 0 12h72a6 6 0 0 0 0-12Z"/>`),
		g.Group(children),
	)
}