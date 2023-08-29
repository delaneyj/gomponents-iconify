package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackPush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.25 7.25L1.75 8L8 11.25L14.25 8l-1.5-.75M1.75 11L8 14.25L14.25 11"/><path d="M8 8.25v-6.5m-2.25 4.5l2.25 2l2.25-2"/></g>`),
		g.Group(children),
	)
}