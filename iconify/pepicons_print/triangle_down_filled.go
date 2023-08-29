package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDownFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13.864 16.52a1 1 0 0 1-1.728 0L6.877 7.503A1 1 0 0 1 7.741 6H18.26a1 1 0 0 1 .864 1.504l-5.26 9.015Z" opacity=".2"/><path d="M10 16a.5.5 0 0 1-.429-.243l-6-10A.5.5 0 0 1 4 5h12a.5.5 0 0 1 .429.757l-6 10A.5.5 0 0 1 10 16Zm5.117-10H4.883L10 14.528L15.117 6Z"/></g>`),
		g.Group(children),
	)
}