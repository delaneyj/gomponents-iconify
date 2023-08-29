package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moneywallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="31" x="4.5" y="8.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.34 16.8a3.6 3.6 0 0 0-3.59 3.6h0a3.6 3.6 0 0 0 3.59 3.6h1.22m0 0h1.22a3.59 3.59 0 0 1 3.59 3.6h0a3.59 3.59 0 0 1-3.59 3.6M24 18c-1-.83-2.06-1.21-4.47-1.21h-1.19M15.09 30c1 .83 2.07 1.21 4.47 1.21h1.22M19.56 33V15m16.79 5.4h7.15v7.2h0h-7.15a2.85 2.85 0 0 1-2.85-2.85v-1.5a2.85 2.85 0 0 1 2.85-2.85Z"/><circle cx="36.47" cy="24" r=".8" fill="currentColor"/>`),
		g.Group(children),
	)
}