package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HashStraightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 148h-44v-40h44a12 12 0 0 0 0-24h-44V40a12 12 0 0 0-24 0v44h-40V40a12 12 0 0 0-24 0v44H40a12 12 0 0 0 0 24h44v40H40a12 12 0 0 0 0 24h44v44a12 12 0 0 0 24 0v-44h40v44a12 12 0 0 0 24 0v-44h44a12 12 0 0 0 0-24Zm-108 0v-40h40v40Z"/>`),
		g.Group(children),
	)
}