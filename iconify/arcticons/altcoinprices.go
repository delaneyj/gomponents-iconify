package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Altcoinprices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.05" d="m33.37 32.57l.12-4.07a9.46 9.46 0 0 0 1-.81C42.08 20 43.77 8.41 41.67 6.34S28 6 20.3 13.59a10.8 10.8 0 0 0-.81 1l-4.07.12A4 4 0 0 0 12 16.75L5.63 28.41a1 1 0 0 0 1.29 1.36l8.58-4.13a13.63 13.63 0 0 0 0 3.72L11.14 35a.71.71 0 0 0 0 .91l.47.44l.44.47a.72.72 0 0 0 .92 0l5.63-4.39a13.63 13.63 0 0 0 3.72 0L18.3 41.1a1 1 0 0 0 1.3 1.28L31.29 36a4 4 0 0 0 2.08-3.38ZM38 10.06a2.06 2.06 0 1 1-2.92 0a2.06 2.06 0 0 1 2.92 0Zm-9.1 9.1a3.44 3.44 0 1 1 4.87 0a3.44 3.44 0 0 1-4.87 0Zm0 0"/>`),
		g.Group(children),
	)
}