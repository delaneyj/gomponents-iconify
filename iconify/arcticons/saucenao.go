package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Saucenao(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42 10.74V7.5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v3.24m0 26.52v3.24a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-3.24M9.5 38.5l29-29m-29 0l29 29"/>`),
		g.Group(children),
	)
}