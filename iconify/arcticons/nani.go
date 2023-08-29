package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nani(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.29 41.94l-7.06-24.78m3.87-5.22l21.28-5.88m-1.54 30.17c2.18-.3 5.71-.35 4.7-3.69S32.4 7.16 32.4 7.16m-9.18 25.15l-3.61-13.12l8.19-2.25l2.74 10.5l-8.12 1.96M12.2 10c-.08 3.68-2.12 14.06-3.92 17"/>`),
		g.Group(children),
	)
}