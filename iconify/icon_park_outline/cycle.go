package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cycle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M13 35H7v6m34 0h-6v-6m0-22h6V7M7 7h6v6"/><path d="M13 7.294C7.578 10.871 4 17.018 4 24c0 1.02.076 2.021.223 3M27 43.776A20.16 20.16 0 0 1 24 44c-6.982 0-13.129-3.578-16.706-9m36.482-14c.148.979.224 1.98.224 3c0 6.982-3.578 13.129-9 16.706M21 4.223A20.16 20.16 0 0 1 24 4c6.982 0 13.129 3.578 16.706 9"/></g>`),
		g.Group(children),
	)
}