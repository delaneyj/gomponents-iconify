package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wolframalpha(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 4.5l5.85 11.85l13.08 1.9l-9.46 9.23l2.23 13.03L24 34.36l-11.7 6.15l2.24-13.03l-9.47-9.23l13.08-1.9L24 4.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 43.5l5.85-11.85l13.08-1.9l-9.46-9.23L35.7 7.49L24 13.64L12.3 7.49l2.24 13.03l-9.47 9.23l13.08 1.9L24 43.5zm-4.09-21.39h8.18m-8.18 4.58h8.18"/>`),
		g.Group(children),
	)
}