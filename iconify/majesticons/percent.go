package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Percent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 18L18 6"/><circle cx="17" cy="17" r="2" fill="currentColor"/><circle cx="7" cy="7" r="2" fill="currentColor"/></g>`),
		g.Group(children),
	)
}