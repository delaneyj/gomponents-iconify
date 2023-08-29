package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Posidonlauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31 8.13V22a7 7 0 0 1-4 6.27V4.5h-6v23.81A7 7 0 0 1 17 22V8.13l-6 3.24V22a13 13 0 0 0 10 12.64v8.86h6v-8.87A13 13 0 0 0 37 22V11.37Z"/>`),
		g.Group(children),
	)
}