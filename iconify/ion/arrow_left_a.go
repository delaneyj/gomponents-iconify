package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M64.5 256.5l192 192v-112h192v-160h-192v-112z" fill="currentColor"/>`),
		g.Group(children),
	)
}