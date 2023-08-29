package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkateboardingLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 17l1.204 1.338A2 2 0 0 0 5.691 19h12.618a2 2 0 0 0 1.487-.662L21 17"/><circle cx="7" cy="21" r="1" fill="currentColor"/><circle cx="17" cy="21" r="1" fill="currentColor"/><circle cx="19" cy="4" r="2" stroke="currentColor" stroke-width="1.5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M15 16.5v-1.25l-.001-.166a3 3 0 0 0-1.493-2.518c-.097-.06-.146-.09-.177-.112a2 2 0 0 1-.184-3.168l.145-.118l.446-.356a1.737 1.737 0 0 0-2.006-2.83L8.5 8M7 15.5h.379c1.358 0 2.66-.54 3.621-1.5m5.5-4a8.246 8.246 0 0 0 4 0"/></g>`),
		g.Group(children),
	)
}