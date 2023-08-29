package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugMinimalisticLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M5 10a7 7 0 0 1 14 0v5a7 7 0 1 1-14 0v-5Z"/><path stroke-linecap="round" d="M19 13h3M5 13H2m18.5-6l-1.798.72M3.5 7l1.798.72M14.5 3.5L17 2M9.5 3.5L7 2m13.5 17l-2-.8m-15 .8l2-.8m5-7.7h3m-3 5h3" opacity=".5"/></g>`),
		g.Group(children),
	)
}