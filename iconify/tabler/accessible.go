package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Accessible(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/><path d="m10 16.5l2-3l2 3m-2-3v-2l3-1m-6 0l3 1"/><circle cx="12" cy="7.5" r=".5" fill="currentColor"/></g>`),
		g.Group(children),
	)
}