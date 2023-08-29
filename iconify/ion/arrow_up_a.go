package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M256.5 64.5l-192 192h112v192h160v-192h112z" fill="currentColor"/>`),
		g.Group(children),
	)
}