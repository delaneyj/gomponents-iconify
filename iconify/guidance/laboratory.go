package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laboratory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M6.5 12.5h.146c2.206 0 4.381.514 6.354 1.5c1.973.986 4.148 1.5 6.354 1.5H20.5m-13-14h9v.25l-.707.972a6.762 6.762 0 0 0 .688 8.759L22.5 17.5c0 2-1 3.5-3 5h-15c-2-1.5-3-3-3-5l6.02-6.02a6.762 6.762 0 0 0 .687-8.758L7.5 1.75V1.5Z"/>`),
		g.Group(children),
	)
}