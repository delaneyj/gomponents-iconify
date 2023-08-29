package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackMediumSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#354a54" d="M53.799 49.21a4.7 4.7 0 0 1-4.701 4.702h-34.2A4.699 4.699 0 0 1 10.2 49.21v-34.2a4.7 4.7 0 0 1 4.698-4.701h34.2a4.7 4.7 0 0 1 4.701 4.701v34.2"/>`),
		g.Group(children),
	)
}