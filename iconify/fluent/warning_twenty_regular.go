package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.562 3.262a.5.5 0 0 1 .88 0l6.5 12a.5.5 0 0 1-.44.739H3.5a.5.5 0 0 1-.44-.738L9.563 3.262Zm1.758-.476c-.567-1.048-2.07-1.048-2.637 0l-6.502 12a1.5 1.5 0 0 0 1.318 2.215h13.003a1.5 1.5 0 0 0 1.319-2.215l-6.5-12ZM10.5 7.5a.5.5 0 0 0-1 0v4a.5.5 0 1 0 1 0v-4Zm.25 6.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Z"/>`),
		g.Group(children),
	)
}