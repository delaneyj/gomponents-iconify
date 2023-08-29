package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FigmaLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 2H8.667a3.333 3.333 0 0 0 0 6.667H12V2Z" opacity=".5"/><path d="M12 8.666H8.667a3.333 3.333 0 0 0 0 6.667H12V8.666Z"/><path d="M18.667 12A3.333 3.333 0 1 1 12 12a3.333 3.333 0 0 1 6.667 0Zm-10 3.334H12v3.333a3.333 3.333 0 1 1-3.333-3.334Z" opacity=".5"/><path d="M12 2h3.333a3.333 3.333 0 1 1 0 6.667H12V2Z"/></g>`),
		g.Group(children),
	)
}