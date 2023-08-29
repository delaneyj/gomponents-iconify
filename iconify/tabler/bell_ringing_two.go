package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellRingingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19.364 4.636a2 2 0 0 1 0 2.828a7 7 0 0 1-1.414 7.072l-2.122 2.12a4 4 0 0 0-.707 3.536L3.808 8.88a4 4 0 0 0 3.535-.707L9.464 6.05a7 7 0 0 1 7.072-1.414a2 2 0 0 1 2.828 0z"/><path d="m7.343 12.414l-.707.707a3 3 0 0 0 4.243 4.243l.707-.707"/></g>`),
		g.Group(children),
	)
}