package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Torchie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.54 11.28L24 5.5l4.44 5.78m0 25.44l-4.46 5.78l-4.44-5.78m22.44-12.71h-7.41m-21.14 0H6.02m10.5 7.47l-5.23 5.24m5.23-20.19l-5.24-5.24m20.18 5.24l5.24-5.24m-5.23 20.19l5.23 5.24"/><circle cx="24" cy="24.01" r="7.37" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}