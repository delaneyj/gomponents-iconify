package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoobarTwoThousand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.36 4.5c-2.46 4.37-5.14 8.27-6.71 15.35a13.68 13.68 0 0 0 0 6.17c.92 5.44 8.69 12.77 14.35 17.47h0c5.66-4.7 13.43-12 14.35-17.47a13.68 13.68 0 0 0 0-6.17c-1.57-7.08-4.25-11-6.71-15.35l-5 6.94h-5.26ZM13.3 18.87c1.7 0 3.95 1.45 5.7 3.63c2.42 3 3.15 6.4 1.64 7.62s-4.7-.21-7.11-3.2s-3.15-6.4-1.64-7.62a2.12 2.12 0 0 1 1.41-.43Zm21.53 0a2.12 2.12 0 0 1 1.41.43c1.51 1.22.78 4.63-1.64 7.62s-5.59 4.42-7.11 3.2s-.78-4.63 1.64-7.62c1.76-2.18 4-3.61 5.7-3.63Zm-15.25 18l3.84-4.32a.72.72 0 0 1 .49-.26H24a.8.8 0 0 1 .59.27l3.84 4.32M24 32.28"/>`),
		g.Group(children),
	)
}