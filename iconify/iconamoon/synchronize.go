package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Synchronize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 0 1 16-5.658"/><path d="M19.5 3v4h-4m5.5 5a9 9 0 0 1-16 5.657"/><path d="M4.5 21v-4h4"/></g>`),
		g.Group(children),
	)
}