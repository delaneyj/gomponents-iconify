package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Venmo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.25 4.45a14.26 14.26 0 0 1 2.06 7.8c0 9.72-8.3 22.34-15 31.2h-15.4L5.74 6.58L19.21 5.3l3.27 26.24c3.05-5 6.81-12.76 6.81-18.08A14.51 14.51 0 0 0 28 6.94Z"/>`),
		g.Group(children),
	)
}