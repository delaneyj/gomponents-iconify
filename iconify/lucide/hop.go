package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.5 5.5C19 7 20.5 9 21 11c-2.5.5-5 .5-8.5-1m-7 7.5C7 19 9 20.5 11 21c.5-2.5.5-5-1-8.5m6.5-1c1 2 1 3.5 1 6c-2.5 0-4 0-6-1m8.5-5c1 1.5 2 3.5 2 4.5c-1.5.5-3 0-4.5-.5m-6 4.5c1.5 1 3.5 2 4.5 2c.5-1.5 0-3-.5-4.5m5-1c1 2 1.5 3.5 1.5 5.5c-2 0-3.5-.5-5.5-1.5"/><path d="M4.783 4.782C8.493 1.072 14.5 1 18 5c-1 1-4.5 2-6.5 1.5c1 1.5 1 4 .5 5.5c-1.5.5-4 .5-5.5-.5C7 13.5 6 17 5 18c-4-3.5-3.927-9.508-.217-13.218ZM4.5 4.5L3 3c-.184-.185-.184-.816 0-1"/></g>`),
		g.Group(children),
	)
}