package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 12a2 2 0 1 0 4 0a2 2 0 0 0-4 0"/><path d="M11.192 17.966C7.95 17.686 5.22 15.697 3 12c2.4-4 5.4-6 9-6c3.326 0 6.14 1.707 8.442 5.122M18.42 15.61a2.1 2.1 0 0 1 2.97 2.97L18 22h-3v-3l3.42-3.39z"/></g>`),
		g.Group(children),
	)
}