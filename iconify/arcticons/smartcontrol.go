package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smartcontrol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.501v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.5 39.501a10.836 10.836 0 0 0 6-2a13.404 13.404 0 0 0 3.844-4.472a25.781 25.781 0 0 0 2.156-5.528c.659-2.34 1.141-4.737 1.988-7.016A13.403 13.403 0 0 1 37.5 14.5a9.43 9.43 0 0 1 6-2m-39 23.001a10.836 10.836 0 0 0 6-2a13.404 13.404 0 0 0 3.844-4.472a25.781 25.781 0 0 0 2.156-5.528c.659-2.34 1.141-4.737 1.988-7.016A13.403 13.403 0 0 1 22.5 10.5a9.43 9.43 0 0 1 6-2"/>`),
		g.Group(children),
	)
}