package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Psiphon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.591 32.866l13.68-.065l3.697-3.483l2.52-20.052L37.3 5.9l-26.332-.32A10.586 10.586 0 0 0 7.51 9.746l10.057.098l-3.964 31.009a13.306 13.306 0 0 0 5.844 1.568Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.795 11.743l-2.173 14.644l9.976-.107l2.138-14.537Z"/>`),
		g.Group(children),
	)
}