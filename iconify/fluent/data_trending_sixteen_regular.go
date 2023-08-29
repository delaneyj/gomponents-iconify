package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTrendingSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 2.5a.5.5 0 0 0-1 0v9A2.5 2.5 0 0 0 4.5 14h9a.5.5 0 0 0 0-1h-9A1.5 1.5 0 0 1 3 11.5v-9Zm7 2a.5.5 0 0 0 .5.5h1.793L9 8.293L7.354 6.646a.5.5 0 0 0-.708 0l-2.5 2.5a.5.5 0 1 0 .708.708L7 7.707l1.646 1.647a.5.5 0 0 0 .708 0L13 5.707v1.87a.5.5 0 0 0 1 0V4.5a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5Z"/>`),
		g.Group(children),
	)
}