package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Casper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.55 4.5a9.29 9.29 0 0 1 9.29 9.29h0a8.9 8.9 0 0 1-.2 1.86L31.18 20h4.59a2.46 2.46 0 0 1 2.46 2.46v1.33a2.46 2.46 0 0 1-2.46 2.46H30.5c0 10 10.63 17.29 10.63 17.29s-21.83.27-26.54-17.29H9.32a2.45 2.45 0 0 1-2.45-2.46v-1.37A2.45 2.45 0 0 1 9.32 20h4.6l-.46-4.26a9.46 9.46 0 0 1-.2-1.91a9.28 9.28 0 0 1 9.29-9.33Z"/>`),
		g.Group(children),
	)
}