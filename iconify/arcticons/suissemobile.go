package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Suissemobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19 5.055h10.293V16.75l10.943-3.488l3.264 9.917l-10.89 3.557l6.942 9.61l-8.412 6.6l-6.977-9.78l-7.215 9.78l-7.934-6.6l6.84-9.61L4.5 23.18l3.386-9.917L19 16.75Z"/>`),
		g.Group(children),
	)
}