package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 4h18"/><rect width="13" height="4" x="6" y="20" fill="currentColor" rx="2" transform="rotate(-90 6 20)"/><rect width="9" height="4" x="14" y="16" fill="currentColor" rx="2" transform="rotate(-90 14 16)"/></g>`),
		g.Group(children),
	)
}