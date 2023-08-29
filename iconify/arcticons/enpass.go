package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Enpass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.513 6.95H10.487a5.985 5.985 0 0 0-5.895 7.02L8.476 36.1a5.985 5.985 0 0 0 5.895 4.95h19.258a5.985 5.985 0 0 0 5.895-4.95l3.884-22.128a5.985 5.985 0 0 0-5.895-7.02Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.527 20.352a6.529 6.529 0 1 0-9.626 5.746v6.957a1.12 1.12 0 0 0 1.12 1.12h3.957a1.12 1.12 0 0 0 1.12-1.12v-6.957a6.526 6.526 0 0 0 3.43-5.745Z"/>`),
		g.Group(children),
	)
}