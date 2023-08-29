package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whatsapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.52 21.52 0 0 0 5.15 34.36L2.5 45.5l11.14-2.65A21.5 21.5 0 1 0 24 2.5Zm-10.75 9.77h5.86a1 1 0 0 1 1 1a10.4 10.4 0 0 0 .66 3.91a1.93 1.93 0 0 1-.66 2.44l-2.05 2a18.6 18.6 0 0 0 3.52 4.79A18.6 18.6 0 0 0 26.35 30l2-2.05c1-1 1.46-1 2.44-.66a10.4 10.4 0 0 0 3.91.66a1.05 1.05 0 0 1 1 1v5.86a1.05 1.05 0 0 1-1 1a23.68 23.68 0 0 1-15.64-6.84a23.6 23.6 0 0 1-6.84-15.64a1.07 1.07 0 0 1 1.03-1.06Z"/>`),
		g.Group(children),
	)
}