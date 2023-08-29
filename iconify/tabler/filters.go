package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filters(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 8a5 5 0 1 0 10 0A5 5 0 1 0 7 8"/><path d="M8 11a5 5 0 1 0 3.998 1.997"/><path d="M12.002 19.003A5 5 0 1 0 16 11"/></g>`),
		g.Group(children),
	)
}