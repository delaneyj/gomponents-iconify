package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSquareInLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M126 136v64a6 6 0 0 1-12 0v-49.51l-69.76 69.75a6 6 0 0 1-8.48-8.48L105.51 142H56a6 6 0 0 1 0-12h64a6 6 0 0 1 6 6Zm82-102H80a14 14 0 0 0-14 14v48a6 6 0 0 0 12 0V48a2 2 0 0 1 2-2h128a2 2 0 0 1 2 2v128a2 2 0 0 1-2 2h-48a6 6 0 0 0 0 12h48a14 14 0 0 0 14-14V48a14 14 0 0 0-14-14Z"/>`),
		g.Group(children),
	)
}