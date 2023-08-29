package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosPlusEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M384 265H264v119h-17V265H128v-17h119V128h17v120h120v17z" fill="currentColor"/>`),
		g.Group(children),
	)
}