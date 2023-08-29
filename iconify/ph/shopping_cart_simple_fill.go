package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartSimpleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M96 216a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm88-16a16 16 0 1 0 16 16a16 16 0 0 0-16-16Zm46.44-132.75A8 8 0 0 0 224 64H48.32l-8.11-28.4A16.08 16.08 0 0 0 24.82 24H8a8 8 0 0 0 0 16h16.82L61 166.59A24.11 24.11 0 0 0 84.07 184h96.11a23.89 23.89 0 0 0 22.94-16.94l28.53-92.71a8 8 0 0 0-1.21-7.1Z"/>`),
		g.Group(children),
	)
}