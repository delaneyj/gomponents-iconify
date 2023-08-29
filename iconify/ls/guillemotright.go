package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Guillemotright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m0 14l505 230l11 7v67l-11 6L0 556v-80l418-192L0 93V14z"/>`),
		g.Group(children),
	)
}