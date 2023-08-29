package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tennis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 4c-.1 6.668-1.738 11.67-4.912 15.004C15.915 22.34 10.885 24.007 4 24.008"/><path stroke-linecap="round" d="M43.968 25.005c-6.512-.447-11.489.946-14.929 4.176c-3.44 3.23-5.118 8.17-5.036 14.819"/></g>`),
		g.Group(children),
	)
}