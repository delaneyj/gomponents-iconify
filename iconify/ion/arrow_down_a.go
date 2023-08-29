package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M256.5 448.5l192-192h-112v-192h-160v192h-112z" fill="currentColor"/>`),
		g.Group(children),
	)
}