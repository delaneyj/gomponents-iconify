package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.001 5h-6v6h6V5Zm2 0v6h6V5h-6Zm6 8h-6v6h6v-6Zm-8 6v-6h-6v6h6Zm-8-16h18v18h-18V3Z"/>`),
		g.Group(children),
	)
}