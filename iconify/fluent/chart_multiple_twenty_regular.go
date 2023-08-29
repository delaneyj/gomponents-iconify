package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartMultipleTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 9.5a6.5 6.5 0 0 0 6 6.481v1.003A7.5 7.5 0 1 1 16.984 9a2 2 0 0 0-1.717 1H10.5A1.5 1.5 0 0 1 9 8.5V3.019A6.5 6.5 0 0 0 3 9.5ZM15.981 9A6.501 6.501 0 0 0 10 3.019V8.5a.5.5 0 0 0 .5.5h5.481ZM16 11a1 1 0 1 1 2 0v7a1 1 0 1 1-2 0v-7Zm-6 4a1 1 0 1 1 2 0v3a1 1 0 1 1-2 0v-3Zm3-2a1 1 0 1 1 2 0v5a1 1 0 1 1-2 0v-5Z"/>`),
		g.Group(children),
	)
}