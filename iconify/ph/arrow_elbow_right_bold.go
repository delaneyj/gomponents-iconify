package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowRightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M244 80v72a12 12 0 0 1-24 0v-43l-91.51 91.52a12 12 0 0 1-17 0l-96-96a12 12 0 0 1 17-17L120 175l83-83h-43a12 12 0 0 1 0-24h72a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}