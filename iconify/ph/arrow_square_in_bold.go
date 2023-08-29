package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSquareInBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M132 136v64a12 12 0 0 1-24 0v-35l-59.51 59.49a12 12 0 0 1-17-17L91 148H56a12 12 0 0 1 0-24h64a12 12 0 0 1 12 12Zm76-108H80a20 20 0 0 0-20 20v44a12 12 0 0 0 24 0V52h120v120h-40a12 12 0 0 0 0 24h44a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Z"/>`),
		g.Group(children),
	)
}