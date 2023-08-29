package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19h18v2H3v-2ZM13 9h7l-8 8l-8-8h7V1h2v8Z"/>`),
		g.Group(children),
	)
}