package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapsArrowDiagonal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.031 8.917l15.477-4.334a.5.5 0 0 1 .616.617l-4.333 15.476a.5.5 0 0 1-.94.067l-3.248-7.382a.5.5 0 0 0-.256-.257L3.965 9.856a.5.5 0 0 1 .066-.94Z"/>`),
		g.Group(children),
	)
}