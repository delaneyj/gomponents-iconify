package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackOverflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M12.412 14.572V10.29h1.428V16H1v-5.71h1.428v4.282h9.984z"/><path d="M3.857 13.145h7.137v-1.428H3.857v1.428zM10.254 0L9.108.852l4.26 5.727l1.146-.852L10.254 0zm-3.54 3.377l5.484 4.567l.913-1.097L7.627 2.28l-.914 1.097zM4.922 6.55l6.47 3.013l.603-1.294l-6.47-3.013l-.603 1.294zm-.925 3.344l6.985 1.469l.294-1.398l-6.985-1.468l-.294 1.397z"/></g>`),
		g.Group(children),
	)
}