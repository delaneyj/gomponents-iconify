package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeavesOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M37 23.879C37 30.573 31.18 36 24 36s-13-5.427-13-12.121C11 17.184 24 4 24 4s13 13.184 13 19.879Z"/><path stroke-linecap="round" d="M24 4v32m0 0v8"/><path d="M37 23.879C37 30.573 31.18 36 24 36s-13-5.427-13-12.12m26-.001C37 17.184 24 4 24 4S11 17.184 11 23.879"/></g>`),
		g.Group(children),
	)
}