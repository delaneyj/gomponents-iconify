package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordMinimalisticLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M21 11.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm-11 0a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Z"/><path stroke-linecap="round" d="M6.5 15h11" opacity=".5"/></g>`),
		g.Group(children),
	)
}