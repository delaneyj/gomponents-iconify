package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhonecallReceiveOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m14.5.5l-6 6m0 0V3m0 3.5H12M2.5.5h2.22a1 1 0 0 1 .97.757l.585 2.345a1 1 0 0 1-.654 1.19l-1.108.37a1.21 1.21 0 0 0-.804 1.385a6.047 6.047 0 0 0 4.744 4.744a1.21 1.21 0 0 0 1.385-.804l.297-.893a1 1 0 0 1 1.396-.578l2.416 1.208a1 1 0 0 1 .553.894V12.5a2 2 0 0 1-2 2h-2c-5.523 0-10-4.477-10-10v-2a2 2 0 0 1 2-2Z"/>`),
		g.Group(children),
	)
}