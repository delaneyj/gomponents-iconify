package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Topic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24a19.94 19.94 0 0 0 3.368 11.112c.244.363-.213 2.66-1.368 6.888c4.229-1.155 6.525-1.612 6.888-1.368A19.94 19.94 0 0 0 24 44Z"/><path stroke-linecap="round" d="M16.605 19.82h16.779m-11.537-4.082l-2.914 16.524m9.914-16.524l-2.914 16.524M14.604 28h16.78"/></g>`),
		g.Group(children),
	)
}