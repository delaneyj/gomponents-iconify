package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistoryThreeBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-dasharray=".5 3.5" d="M17 3.338A9.954 9.954 0 0 0 12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10a9.966 9.966 0 0 0-.832-4"/><path d="M22 12c0-1.821-.487-3.53-1.338-5M12 2c1.821 0 3.53.487 5 1.338"/><path stroke-linejoin="round" d="M12 9v4h4"/></g>`),
		g.Group(children),
	)
}