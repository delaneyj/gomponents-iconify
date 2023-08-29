package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diaspora(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19 4.49h10v13.08l12.45-4L44.5 23l-12.45 4l7.67 10.57l-8 5.87L24 32.94l-7.67 10.57l-8-5.87L16 27.07L3.5 23l3.09-9.5l12.45 4Z"/>`),
		g.Group(children),
	)
}