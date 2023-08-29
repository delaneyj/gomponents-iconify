package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nayatv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.582 23.95V43M5.708 23.95h19.119a1.282 1.282 0 0 1 1.21.884l5.748 17.122l9.635-25.861"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.926 43H8.074a2.366 2.366 0 0 1-2.366-2.366v-22.83a1.94 1.94 0 0 1 1.94-1.94h32.278a2.366 2.366 0 0 1 2.366 2.367v22.404A2.366 2.366 0 0 1 39.926 43ZM16.38 13.316a3.587 3.587 0 0 1 7.163 0m-1.444-3.063L26.176 5m-8.1 5.253L15.343 7.09"/>`),
		g.Group(children),
	)
}