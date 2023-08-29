package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarOverdueTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M21 12.022V8.5H3v9.25A3.25 3.25 0 0 0 6.25 21h5.772A6.5 6.5 0 0 1 21 12.022z" fill="currentColor"/><path d="M21 6.25A3.25 3.25 0 0 0 17.75 3H6.25A3.25 3.25 0 0 0 3 6.25V7h18v-.75z" fill="currentColor"/><path d="M23 17.5a5.5 5.5 0 1 0-11 0a5.5 5.5 0 0 0 11 0zm-6-3a.5.5 0 0 1 1 0v4a.5.5 0 0 1-1 0v-4zm1.125 6a.625.625 0 1 1-1.25 0a.625.625 0 0 1 1.25 0z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}