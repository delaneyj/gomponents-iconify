package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 2H1.5a.5.5 0 0 0-.5.5V7h6V2Zm1 0v5h6V2.5a.5.5 0 0 0-.5-.5H8ZM7 8H1v4.5a.5.5 0 0 0 .5.5H7V8Zm1 5V8h6v4.5a.5.5 0 0 1-.5.5H8ZM1.5 1A1.5 1.5 0 0 0 0 2.5v10A1.5 1.5 0 0 0 1.5 14h12a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 13.5 1h-12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}