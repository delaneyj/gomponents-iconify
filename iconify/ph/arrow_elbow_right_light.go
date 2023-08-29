package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowRightLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M238 80v72a6 6 0 0 1-12 0V94.48L124.24 196.24a6 6 0 0 1-8.48 0l-96-96a6 6 0 0 1 8.48-8.48L120 183.51L217.52 86H160a6 6 0 0 1 0-12h72a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}