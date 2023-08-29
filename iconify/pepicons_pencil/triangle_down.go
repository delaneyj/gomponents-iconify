package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 16a.5.5 0 0 1-.429-.243l-6-10A.5.5 0 0 1 4 5h12a.5.5 0 0 1 .429.757l-6 10A.5.5 0 0 1 10 16Zm5.117-10H4.883L10 14.528L15.117 6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}