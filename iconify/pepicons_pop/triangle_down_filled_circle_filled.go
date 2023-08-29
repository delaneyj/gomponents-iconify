package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDownFilledCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopTriangleDownFilledCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M13.866 18.499a1 1 0 0 1-1.732 0l-5.769-10A1 1 0 0 1 7.231 7H18.77a1 1 0 0 1 .866 1.5l-5.769 9.999Z"/><path d="M6.365 8.5A1 1 0 0 1 7.231 7H18.77a1 1 0 0 1 .866 1.5l-5.769 9.999a1 1 0 0 1-1.732 0l-5.769-10ZM13 13.997L15.306 10h-4.612L13 13.997Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopTriangleDownFilledCircleFilled0)"/></g>`),
		g.Group(children),
	)
}