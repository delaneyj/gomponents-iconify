package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyOppo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.177 16.018l-7.37-7.37H12.193l-7.37 7.37a1.102 1.102 0 0 0-.058 1.496l18.399 21.454a1.101 1.101 0 0 0 1.672 0l18.399-21.454a1.102 1.102 0 0 0-.058-1.496Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.673 20.336L24 27.664l7.327-7.328"/>`),
		g.Group(children),
	)
}