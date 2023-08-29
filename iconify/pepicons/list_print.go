package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><g opacity=".8"><circle cx="6" cy="7.25" r="1.5" fill="currentColor"/><circle cx="6" cy="11.25" r="1.5" fill="currentColor"/><circle cx="6" cy="15.25" r="1.5" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M9.5 7.25h7m-7 4h7m-7 4h7"/></g><circle cx="5" cy="6" r="1.5" fill="currentColor"/><circle cx="5" cy="10" r="1.5" fill="currentColor"/><circle cx="5" cy="14" r="1.5" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" d="M8 6.5h8m-8 4h8m-8 4h8"/></g>`),
		g.Group(children),
	)
}