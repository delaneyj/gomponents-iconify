package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Satstat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.19 6.43h.17a.91.91 0 0 1 .53.32L34.61 28a.91.91 0 0 1-.34 1.38A17.77 17.77 0 0 1 28 30.66a14 14 0 0 1-11.65-5.48a14 14 0 0 1-2.62-12.61a17.92 17.92 0 0 1 2.73-5.79a.91.91 0 0 1 .72-.35Zm16.17 2a2.7 2.7 0 1 1-1.47.57a2.74 2.74 0 0 1 1.47-.57ZM18.93 32.1h8.34a1.08 1.08 0 0 1 .31 0a1.67 1.67 0 0 1 1.24 1l2.28 5.61a1.69 1.69 0 0 1-1.56 2.31H16.66a1.68 1.68 0 0 1-1.55-2.31l2.27-5.61a1.66 1.66 0 0 1 1.55-1Zm.19-26.17l10 3.57m4.92 6.28l1.17 10.51"/>`),
		g.Group(children),
	)
}