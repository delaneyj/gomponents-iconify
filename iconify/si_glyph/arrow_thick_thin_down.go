package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickThinDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M2.336 8.667c.199 0 .404.045.596.141l6.07 3.034l6.069-3.034a1.332 1.332 0 1 1 1.192 2.385l-6.666 3.333a1.332 1.332 0 0 1-1.193 0l-6.666-3.333a1.332 1.332 0 0 1 .598-2.526z"/><path d="M2.336 4.334c.1 0 .201.022.297.07l6.369 3.184l6.367-3.184a.669.669 0 0 1 .895.298a.668.668 0 0 1-.299.895L9.299 8.93a.665.665 0 0 1-.596 0L2.037 5.597a.668.668 0 0 1 .299-1.263z"/></g>`),
		g.Group(children),
	)
}