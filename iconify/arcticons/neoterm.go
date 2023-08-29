package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Neoterm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.204 30.247A4.475 4.475 0 0 0 16.127 32h2.369a3.996 3.996 0 0 0 3.99-4h0a3.996 3.996 0 0 0-3.99-4h-2.618a3.996 3.996 0 0 1-3.991-4h0a3.996 3.996 0 0 1 3.991-4h2.37a4.476 4.476 0 0 1 3.922 1.753M17.187 16v-2m0 20v-2m8.326 2h10.6"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}