package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 4v2h-1l-5 7.5V22H9v-8.5L4 6H3V4h18ZM6.404 6L11 12.894V20h2v-7.106L17.596 6H6.404Z"/>`),
		g.Group(children),
	)
}