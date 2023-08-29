package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volkswagen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5c11.876 0 21.5 9.624 21.5 21.5S35.876 45.5 24 45.5S2.5 35.876 2.5 24S12.124 2.5 24 2.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.297 16.894L31.51 40.74L24 25.386L16.49 40.74L3.703 16.894"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.825 4.39L24 22.47L15.174 4.39"/>`),
		g.Group(children),
	)
}