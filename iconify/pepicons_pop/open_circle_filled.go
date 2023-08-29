package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopOpenCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M13.707 13.707a1 1 0 0 1-1.414-1.414l6-6a1 1 0 1 1 1.414 1.414l-6 6Z"/><path d="M18 18v-3.5a1 1 0 1 1 2 0V19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h4.5a1 1 0 0 1 0 2H8v10h10Zm2-7a1 1 0 1 1-2 0V7a1 1 0 1 1 2 0v4Z"/><path d="M15 8a1 1 0 1 1 0-2h4a1 1 0 1 1 0 2h-4Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopOpenCircleFilled0)"/></g>`),
		g.Group(children),
	)
}