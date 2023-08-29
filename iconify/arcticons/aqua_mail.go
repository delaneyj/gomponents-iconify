package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AquaMail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.986 40.303l12.203-9.238a3 3 0 0 1 3.622 0l12.2 9.243m-17-8.352L3.496 17.517m23.503 14.449L44.51 17.53m-7.84.417l-10.78 8.872a3.104 3.104 0 0 1-3.947-.002l-10.77-8.879c9.215-2.826 14.002 5.994 25.496.009Z"/>`),
		g.Group(children),
	)
}