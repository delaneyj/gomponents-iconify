package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Intercom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M13 14a3 3 0 0 1 3-3h16a3 3 0 0 1 3 3v11l-2 6v10a3 3 0 0 1-3 3H18a3 3 0 0 1-3-3V31l-2-6V14Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M19 11V4m9 7V7m0 12h-8m6 7h-4"/></g>`),
		g.Group(children),
	)
}