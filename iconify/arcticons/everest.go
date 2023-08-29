package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Everest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 40.021L19.375 7.98l6.373 11.513l3.398-6.25L43.5 39.993l-39 .03Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.355 34.845l7.02-15.81l6.373 9.904l3.398-6.234l6.69 12.267l-23.481-.127Z"/>`),
		g.Group(children),
	)
}