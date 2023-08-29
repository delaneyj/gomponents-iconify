package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 17v2H11v-2h11M19 4.5v5h3L16.5 15L11 9.5h3v-5h5M10.7 15H8.8l-.7-2H4.9l-.7 2H2.3l3.2-9h2l3.2 9m-3.05-3.35L6.5 8l-1.15 3.65h2.3Z"/>`),
		g.Group(children),
	)
}