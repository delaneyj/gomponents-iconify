package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.84 12.663v9.39L0 20.697v-8.034zm0-10.72v9.505H0V3.303zM24 12.663V24l-13.082-1.803v-9.534zM24 0v11.452H10.918V1.803z"/>`),
		g.Group(children),
	)
}