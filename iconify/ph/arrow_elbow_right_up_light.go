package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowRightUpLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228.24 100.24a6 6 0 0 1-8.48 0L182 62.49V192a6 6 0 0 1-6 6H32a6 6 0 0 1 0-12h138V62.49l-37.76 37.75a6 6 0 1 1-8.48-8.48l48-48a6 6 0 0 1 8.48 0l48 48a6 6 0 0 1 0 8.48Z"/>`),
		g.Group(children),
	)
}