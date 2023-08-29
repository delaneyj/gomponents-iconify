package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adaptivebrightnesstile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.9 24.87H27m1.56 4.67L24 15.69l-4.66 13.85M24 6l5.54 5.54h6.92v6.92L42 24l-5.54 5.54v6.92h-6.92L24 42l-5.54-5.54h-6.92v-6.92L6 24l5.54-5.54v-6.92h6.92Z"/>`),
		g.Group(children),
	)
}