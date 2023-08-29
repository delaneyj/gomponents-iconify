package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notificationcron(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m38.94 42.18l-1.27-9.61l3.16-8.65A12.67 12.67 0 0 0 36.09 9.1a2.22 2.22 0 0 0 .18-.38a3.21 3.21 0 0 0-6-2.19a2.62 2.62 0 0 0-.11.4A12.7 12.7 0 0 0 17 15.2l-3.16 8.72l-7.18 6.54a.77.77 0 0 0-.06 1.11a.71.71 0 0 0 .32.22l3.56 1.29h0l16.61 6h0a3.22 3.22 0 0 0 6 2.37l.06-.17h0l1.17.43h0L37.86 43a.79.79 0 0 0 1-.44a.78.78 0 0 0 .08-.38ZM27.09 39.1l6.04 2.2"/>`),
		g.Group(children),
	)
}