package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderForbidFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 11.255A7 7 0 0 0 12.255 21H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h7.414l2 2H21a1 1 0 0 1 1 1v5.255ZM18 22a5 5 0 1 1 0-10a5 5 0 0 1 0 10Zm-1.293-2.292a3 3 0 0 0 4.001-4.001l-4.001 4Zm-1.415-1.415l4.001-4a3 3 0 0 0-4.001 4.001Z"/>`),
		g.Group(children),
	)
}