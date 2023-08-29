package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.529 4.18a9.5 9.5 0 0 0-3.76-.677a.75.75 0 1 1-.037-1.5A11 11 0 0 1 13.997 13.27a.75.75 0 1 1-1.5-.036A9.5 9.5 0 0 0 6.53 4.18ZM4.624 8.803a4.5 4.5 0 0 0-1.838-.298a.75.75 0 1 1-.07-1.498a6 6 0 0 1 6.277 6.282a.75.75 0 0 1-1.498-.072a4.5 4.5 0 0 0-2.871-4.414ZM3 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}