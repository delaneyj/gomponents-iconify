package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignRightLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M34 64a6 6 0 0 1 6-6h176a6 6 0 0 1 0 12H40a6 6 0 0 1-6-6Zm182 34H88a6 6 0 0 0 0 12h128a6 6 0 0 0 0-12Zm0 40H40a6 6 0 0 0 0 12h176a6 6 0 0 0 0-12Zm0 40H88a6 6 0 0 0 0 12h128a6 6 0 0 0 0-12Z"/>`),
		g.Group(children),
	)
}