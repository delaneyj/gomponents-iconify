package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDCollection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M16.56 19.43L24 14.8l7.44 4.628"/><path d="M41.007 14.801L24 4.5L6.993 14.801M24 43.5l17.007-10.098v-18.6l-9.567 4.627v8.659l-5.846 3.188v5.315L24 37.594l-1.594-1.003v-5.315l-5.846-3.188v-8.659l-9.567-4.628v18.601L24 43.5Z"/></g>`),
		g.Group(children),
	)
}