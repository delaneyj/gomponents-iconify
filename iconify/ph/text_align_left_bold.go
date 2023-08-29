package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M28 64a12 12 0 0 1 12-12h176a12 12 0 0 1 0 24H40a12 12 0 0 1-12-12Zm12 52h128a12 12 0 0 0 0-24H40a12 12 0 0 0 0 24Zm176 16H40a12 12 0 0 0 0 24h176a12 12 0 0 0 0-24Zm-48 40H40a12 12 0 0 0 0 24h128a12 12 0 0 0 0-24Z"/>`),
		g.Group(children),
	)
}