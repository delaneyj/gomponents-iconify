package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackPop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.25 6.75L1.75 8L8 11.25L14.25 8l-2.5-1.25M1.75 11L8 14.25L14.25 11M8 8.25v-6.5m-2.25 2l2.25-2l2.25 2"/>`),
		g.Group(children),
	)
}