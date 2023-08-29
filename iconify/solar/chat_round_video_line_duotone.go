package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatRoundVideoLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.6.376 3.112 1.043 4.453c.178.356.237.763.134 1.148l-.595 2.226a1.3 1.3 0 0 0 1.591 1.591l2.226-.595a1.634 1.634 0 0 1 1.149.133A9.958 9.958 0 0 0 12 22Z"/><path stroke-linecap="round" d="M16 12c0-.845-.927-1.414-2.78-2.553c-1.88-1.155-2.82-1.732-3.52-1.308C9 8.563 9 9.709 9 12c0 2.291 0 3.437.7 3.861c.701.424 1.64-.153 3.52-1.308C15.072 13.414 16 12.844 16 12Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}