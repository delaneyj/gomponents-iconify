package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M13 19a8 8 0 0 0-8-8m14 8c0-7.732-6.268-14-14-14"/><circle cx="6" cy="18" r="2" fill="currentColor"/></g>`),
		g.Group(children),
	)
}