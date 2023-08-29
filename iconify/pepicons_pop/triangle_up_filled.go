package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleUpFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M9.134 4.501a1 1 0 0 1 1.732 0l5.769 10A1 1 0 0 1 15.769 16H4.23a1 1 0 0 1-.866-1.5l5.769-9.999Z"/><path d="M16.635 14.5a1 1 0 0 1-.866 1.5H4.23a1 1 0 0 1-.866-1.5l5.769-9.999a1 1 0 0 1 1.732 0l5.769 10ZM10 9.003L7.694 13h4.612L10 9.003Z"/></g>`),
		g.Group(children),
	)
}