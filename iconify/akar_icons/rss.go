package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 10.418c6.068-.319 9.9 3.514 9.582 9.582"/><circle cx="5" cy="19" r="1"/><path d="M4 4.03C14.114 3.5 20.501 9.887 19.97 20"/></g>`),
		g.Group(children),
	)
}