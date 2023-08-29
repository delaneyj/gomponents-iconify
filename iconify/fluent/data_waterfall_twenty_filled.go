package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataWaterfallTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 3.5a.5.5 0 0 1 .5-.5H7a2 2 0 0 1 2 2v4.5h6c1.035 0 2 .741 2 1.8V16h1.5a.5.5 0 0 1 0 1H13c-1.035 0-2-.741-2-1.8v-4.7H5a2 2 0 0 1-2-2V4H1.5a.5.5 0 0 1-.5-.5Z"/>`),
		g.Group(children),
	)
}