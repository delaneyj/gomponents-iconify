package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func User(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="8" r="5" fill="currentColor"/><path d="M20 21a8 8 0 1 0-16 0"/><path fill="currentColor" d="M12 13a8 8 0 0 0-8 8h16a8 8 0 0 0-8-8z"/></g>`),
		g.Group(children),
	)
}