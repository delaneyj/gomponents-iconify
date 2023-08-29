package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTrendingTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 3.5a.5.5 0 0 0-1 0v11A2.5 2.5 0 0 0 5.5 17h11a.5.5 0 0 0 0-1h-11A1.5 1.5 0 0 1 4 14.5v-11ZM12.5 5a.5.5 0 0 0 0 1h2.793L11 10.293L9.354 8.646a.5.5 0 0 0-.708 0l-3.5 3.5a.5.5 0 0 0 .708.708L9 9.707l1.646 1.647a.5.5 0 0 0 .707 0L16 6.707V9.5a.5.5 0 0 0 1 0v-4a.5.5 0 0 0-.5-.5h-4Z"/>`),
		g.Group(children),
	)
}