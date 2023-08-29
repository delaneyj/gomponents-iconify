package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAsc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19 3l4 5h-3v12h-2V8h-3l4-5Zm-5 15v2H3v-2h11Zm0-7v2H3v-2h11Zm-2-7v2H3V4h9Z"/>`),
		g.Group(children),
	)
}