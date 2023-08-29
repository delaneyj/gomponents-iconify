package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hotdeath(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.54 32.79l7.6 3.94l-1.22-25.89a54.25 54.25 0 0 1 2.72 7.41a41.05 41.05 0 0 1 1 13.42a6.69 6.69 0 0 1-1 3.19a4.6 4.6 0 0 1-1.22.84l2.16.66l2.34 5.16M25.26 5.4s7.15-1.57 10.5-.56a9 9 0 0 1 4.79 3.56c1.68 2.58 1.74 5.95 2.06 9a60.32 60.32 0 0 1 .28 12.38a25.51 25.51 0 0 1-2.44 8.73a8.7 8.7 0 0 1-2.53 3a7.87 7.87 0 0 1-2.81 1.31c-3.19.82-9.85.66-9.85.66l-.94-9.1l1.22-1.6l-.75-13l1-1L24.6 7.65ZM4.9 11l.8 13l-.7 2.79v16.32h7.22l.47-16.41l2.35 2.9l-.61 13.32l8 .19l.18-35.17l-7.31-.1l-.47 13.6l-2.16-.18l.02-13.14h-6Z"/>`),
		g.Group(children),
	)
}