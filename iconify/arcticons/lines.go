package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lines(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 19.446v9.108h4.554m2.562-9.108v9.108m11.159 0h4.554m-4.554-9.108h4.554M26.775 24h2.969m-2.969-4.554v9.108m-8.596 0v-9.108l6.033 9.108v-9.108m9.435 8.11c.558.727 1.259.998 2.233.998h1.348a2.274 2.274 0 0 0 2.272-2.277h0A2.274 2.274 0 0 0 37.228 24h-1.49a2.274 2.274 0 0 1-2.272-2.277h0a2.274 2.274 0 0 1 2.272-2.277h1.348c.975 0 1.675.27 2.234.998"/>`),
		g.Group(children),
	)
}