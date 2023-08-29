package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.586 2H11a1 1 0 0 1 0-2h4a1 1 0 0 1 1 1v4a1 1 0 0 1-2 0V3.414L9.414 8L14 12.586V11a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1h-4a1 1 0 0 1 0-2h1.586L8 9.414L3.414 14H5a1 1 0 0 1 0 2H1a1 1 0 0 1-1-1v-4a1 1 0 0 1 2 0v1.586L6.586 8L2 3.414V5a1 1 0 1 1-2 0V1a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2H3.414L8 6.586L12.586 2z"/>`),
		g.Group(children),
	)
}