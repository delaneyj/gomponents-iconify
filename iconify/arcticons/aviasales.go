package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aviasales(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.34 31.4a5 5 0 0 1-5-5v-4.78a5 5 0 0 1 5-5h1.16V7.5a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v9.12h1.16a5 5 0 0 1 5 5v4.78a5 5 0 0 1-5 5H5.5v9.1a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-9.1Z"/><rect width="18.05" height="23.92" x="14.99" y="13.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6"/><rect width="12.79" height="16.04" x="17.62" y="15.7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.57 27v1.27a1.5 1.5 0 0 0 1.5 1.5h1.27"/>`),
		g.Group(children),
	)
}