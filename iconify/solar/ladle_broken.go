package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LadleBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M2 5.684a3.684 3.684 0 0 1 7.368 0M22 14.5v1.184a6.316 6.316 0 0 1-12.632 0V9.158"/><path d="M16 12c-3.054 0-6.5 1.12-6.5 2.5S12.946 17 16 17s6-1.12 6-2.5c0-.72-.8-1.369-2-1.825"/></g>`),
		g.Group(children),
	)
}