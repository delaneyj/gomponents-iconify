package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomenBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><circle cx="12" cy="9" r="7" opacity=".5"/><path d="M11.25 15.96v1.79H9.5a.75.75 0 0 0 0 1.5h1.75V22a.75.75 0 0 0 1.5 0v-2.75h1.75a.75.75 0 0 0 0-1.5h-1.75v-1.79a7.091 7.091 0 0 1-1.5 0Z"/></g>`),
		g.Group(children),
	)
}