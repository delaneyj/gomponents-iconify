package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistoryTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M2 12c0 5.523 4.477 10 10 10c1.821 0 3.53-.487 5-1.338M12 2c5.523 0 10 4.477 10 10c0 1.821-.487 3.53-1.338 5"/><path stroke-linejoin="round" d="M12 9v4h4"/><path stroke-dasharray=".5 3.5" d="M17 20.662A9.955 9.955 0 0 1 12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c0 1.821-.487 3.53-1.338 5"/></g>`),
		g.Group(children),
	)
}