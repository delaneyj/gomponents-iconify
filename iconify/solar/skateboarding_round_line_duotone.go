package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkateboardingRoundLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m3 17l.608.676A4 4 0 0 0 6.581 19H17.42a4 4 0 0 0 2.973-1.324L21 17"/><circle cx="7" cy="21" r="1" fill="currentColor" opacity=".5"/><circle cx="17" cy="21" r="1" fill="currentColor" opacity=".5"/><circle cx="19" cy="4" r="2" stroke="currentColor" stroke-width="1.5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M15 16.5v-2.133a1.85 1.85 0 0 0-.666-1.422l-.996-.83a1.59 1.59 0 0 1-.106-2.346l1.654-1.654a1.067 1.067 0 0 0-.335-1.736a4.269 4.269 0 0 0-3.944.304L8.5 8"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m11 14l-.621.621c-.434.434-.65.65-.926.765c-.276.114-.583.114-1.196.114H7m9.5-5.5h3" opacity=".5"/></g>`),
		g.Group(children),
	)
}