package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M11.5 4.5A1.5 1.5 0 0 0 10 6v3.5H6.5a1.5 1.5 0 0 0 0 3H10V16a1.5 1.5 0 0 0 3 0v-3.5h3.5a1.5 1.5 0 0 0 0-3H13V6a1.5 1.5 0 0 0-1.5-1.5Z" clip-rule="evenodd" opacity=".8"/><path d="M5 10.5a.5.5 0 0 1 0-1h10a.5.5 0 0 1 0 1H5Z"/><path d="M9.5 5a.5.5 0 0 1 1 0v10a.5.5 0 0 1-1 0V5Z"/></g>`),
		g.Group(children),
	)
}