package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.791 28.976l-3.756-9.952l-3.755 9.952m15.783-9.952l-3.756 9.952l-3.756-9.952h7.512zM8.5 22.78a3.756 3.756 0 0 1 3.756-3.756h0m-3.756 0v9.952m30.508-1.895a3.755 3.755 0 0 1-3.264 1.895h0a3.756 3.756 0 0 1-3.756-3.755v-2.442a3.756 3.756 0 0 1 3.756-3.755h0a3.756 3.756 0 0 1 3.756 3.755V24h-7.512"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}