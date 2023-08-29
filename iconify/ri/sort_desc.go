package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortDesc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 4v12h3l-4 5l-4-5h3V4h2Zm-8 14v2H3v-2h9Zm2-7v2H3v-2h11Zm0-7v2H3V4h11Z"/>`),
		g.Group(children),
	)
}