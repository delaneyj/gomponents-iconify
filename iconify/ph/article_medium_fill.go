package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArticleMediumFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M56 144a8 8 0 0 1-8 8H24a8 8 0 0 1 0-16h8V72h-8a8 8 0 0 1 0-16h16a8 8 0 0 1 6.78 3.74L80 112.91l33.22-53.15A8 8 0 0 1 120 56h16a8 8 0 0 1 0 16h-8v64h8a8 8 0 0 1 0 16h-24a8 8 0 0 1 0-16V91.89l-25.22 40.35a8 8 0 0 1-13.56 0L48 91.89V136a8 8 0 0 1 8 8Zm184-24h-72a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8h72a8 8 0 0 0 8-8v-16a8 8 0 0 0-8-8Zm0 48H72a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8h168a8 8 0 0 0 8-8v-16a8 8 0 0 0-8-8Zm0-96h-72a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8h72a8 8 0 0 0 8-8V80a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}