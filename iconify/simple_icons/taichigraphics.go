package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taichigraphics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.343 20.672a1.94 1.94 0 0 0 1.94-1.94a1.94 1.94 0 1 0-3.88 0a1.94 1.94 0 0 0 1.94 1.94zM9.058 12.796a6.858 6.858 0 1 0-2.324-9.67c-.062.099-.125.198-.185.3c-.06.103-.11.205-.167.31a6.858 6.858 0 0 0 2.676 9.06zm0-.003h-4.23l-2.113 3.663l2.114 3.667h4.229l2.116-3.667zm0 7.33L6.82 23.999h4.48Z"/>`),
		g.Group(children),
	)
}