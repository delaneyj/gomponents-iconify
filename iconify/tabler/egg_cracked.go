package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EggCracked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 14.083c0 4.154-2.966 6.74-7 6.917c-4.2 0-7-2.763-7-6.917C5 8.545 8.5 2.993 12 3c3.5.007 7 5.545 7 11.083z"/><path d="m12 3l-1.5 5l3.5 2.5l-2 3.5"/></g>`),
		g.Group(children),
	)
}