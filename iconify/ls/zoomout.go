package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zoomout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m541 517l176 176l-72 73l-184-185c-45 25-97 39-151 39C139 620 0 481 0 310S139 0 310 0s311 139 311 310c0 80-31 153-80 207zM86 310c0 124 100 223 224 223c123 0 224-99 224-223S433 87 310 87C186 87 86 186 86 310zm56-35h337v71H142v-71z"/>`),
		g.Group(children),
	)
}