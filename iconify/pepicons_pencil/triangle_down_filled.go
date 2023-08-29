package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDownFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M10.866 15.499a1 1 0 0 1-1.732 0l-5.769-10A1 1 0 0 1 4.231 4H15.77a1 1 0 0 1 .866 1.5l-5.769 9.999Z"/><path d="M3.365 5.5A1 1 0 0 1 4.231 4H15.77a1 1 0 0 1 .866 1.5l-5.769 9.999a1 1 0 0 1-1.732 0l-5.769-10ZM10 10.997L12.306 7H7.694L10 10.997Z"/></g>`),
		g.Group(children),
	)
}