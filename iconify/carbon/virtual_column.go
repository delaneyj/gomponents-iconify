package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirtualColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m25 11l1.414 1.414L23.828 15H30v2h-6.172l2.586 2.586L25 21l-5-5l5-5z"/><path fill="currentColor" d="M30 28H20a2.002 2.002 0 0 1-2-2V6a2.002 2.002 0 0 1 2-2h10v2H20v20h10zm-18 0H2v-2h10V6H2V4h10a2.002 2.002 0 0 1 2 2v20a2.002 2.002 0 0 1-2 2z"/><path fill="currentColor" d="m7 11l-1.414 1.414L8.172 15H2v2h6.172l-2.586 2.586L7 21l5-5l-5-5z"/>`),
		g.Group(children),
	)
}