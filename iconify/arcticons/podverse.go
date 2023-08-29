package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Podverse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.21 25.17v17.02m23.58-17.02v17.02M8.88 22.55A15.12 15.12 0 0 1 24 7.43h0a15.12 15.12 0 0 1 15.12 15.12v4.53a8.81 8.81 0 0 0-3.33-1.91a1.31 1.31 0 0 0-1.31-1.31h-2.62a1.3 1.3 0 0 0-1.31 1.31v17a1.3 1.3 0 0 0 1.31 1.31h2.62a1.31 1.31 0 0 0 1.31-1.31a8.92 8.92 0 0 0 6.27-8.51V22.55A18.06 18.06 0 0 0 24 4.5h0A18.06 18.06 0 0 0 5.94 22.55v11.13a8.92 8.92 0 0 0 6.27 8.51a1.31 1.31 0 0 0 1.31 1.31h2.62a1.3 1.3 0 0 0 1.31-1.31v-17a1.3 1.3 0 0 0-1.31-1.31h-2.62a1.31 1.31 0 0 0-1.31 1.31a8.81 8.81 0 0 0-3.33 1.91v-4.55Z"/>`),
		g.Group(children),
	)
}