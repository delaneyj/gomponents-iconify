package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendLeftDownBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 32a12 12 0 0 1-12 12a84.09 84.09 0 0 0-84 84v67l27.51-27.52a12 12 0 0 1 17 17l-48 48a12 12 0 0 1-17 0l-48-48a12 12 0 0 1 17-17L92 195v-67A108.12 108.12 0 0 1 200 20a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}