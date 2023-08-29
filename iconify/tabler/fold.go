package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v6l3-3M9 6l3 3m0 12v-6l3 3m-6 0l3-3m-8-3h1m4 0h1m4 0h1m4 0h1"/>`),
		g.Group(children),
	)
}