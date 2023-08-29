package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBottomLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 20h18"/><rect width="13" height="4" x="6" y="17" rx="2" transform="rotate(-90 6 17)"/><rect width="9" height="4" x="14" y="17" rx="2" transform="rotate(-90 14 17)"/></g>`),
		g.Group(children),
	)
}