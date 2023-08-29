package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowMerge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m243 371l-72-72l30-30l72 72zM41 107l96-96l96 96h-75v136L30 371L0 341l115-115V107H41z"/>`),
		g.Group(children),
	)
}