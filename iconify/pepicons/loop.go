package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M4.475 4.475a5.5 5.5 0 1 0 7.778 7.778a5.5 5.5 0 0 0-7.778-7.778Zm6.364 6.364a3.5 3.5 0 1 1-4.95-4.95a3.5 3.5 0 0 1 4.95 4.95Z" clip-rule="evenodd"/><path d="M11.192 13.314a1.5 1.5 0 1 1 2.122-2.122l3.535 3.536a1.5 1.5 0 1 1-2.121 2.121l-3.536-3.535Z"/></g>`),
		g.Group(children),
	)
}