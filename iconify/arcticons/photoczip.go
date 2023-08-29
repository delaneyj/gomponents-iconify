package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Photoczip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 33V15h5.9a6.05 6.05 0 0 1 0 12.1h-5.9M31.53 24h5.97l-5.97 9.01h5.97M29.4 30v0a3 3 0 0 1-3 3h0a3 3 0 0 1-3-3v-3a3 3 0 0 1 3-3h0a3 3 0 0 1 3 3v0"/>`),
		g.Group(children),
	)
}