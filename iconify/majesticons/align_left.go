package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 3v18"/><rect width="13" height="4" x="7" y="6" fill="currentColor" rx="2"/><rect width="9" height="4" x="7" y="14" fill="currentColor" rx="2"/></g>`),
		g.Group(children),
	)
}