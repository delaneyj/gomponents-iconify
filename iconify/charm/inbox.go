package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 13.25h12.5v-5l-2.5-5.5h-7.5l-2.5 5.5z"/><path d="M2.25 8.75h3l1 1.5h3.5l1-1.5h3"/></g>`),
		g.Group(children),
	)
}