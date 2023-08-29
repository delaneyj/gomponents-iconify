package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTopSimpleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M168 72v152a8 8 0 0 1-8 8H96a8 8 0 0 1-8-8V72a8 8 0 0 1 8-8h64a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M208 32a8 8 0 0 1-8 8H56a8 8 0 0 1 0-16h144a8 8 0 0 1 8 8Zm-32 40v152a16 16 0 0 1-16 16H96a16 16 0 0 1-16-16V72a16 16 0 0 1 16-16h64a16 16 0 0 1 16 16Zm-16 0H96v152h64Z"/></g>`),
		g.Group(children),
	)
}