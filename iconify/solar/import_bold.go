package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImportBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M15.53 10.47a.75.75 0 0 0-1.06 0l-1.72 1.72V4a.75.75 0 0 0-1.5 0v8.19l-1.72-1.72a.75.75 0 0 0-1.06 1.06l3 3a.75.75 0 0 0 1.06 0l3-3a.75.75 0 0 0 0-1.06Z" clip-rule="evenodd"/><path d="M17.748 12c-.448 0-.84.274-1.157.591l-3 3a2.25 2.25 0 0 1-3.182 0l-3-3C7.092 12.274 6.7 12 6.252 12H4a8 8 0 1 0 16 0h-2.252Z"/></g>`),
		g.Group(children),
	)
}