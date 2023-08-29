package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatSmileLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.455 19L2 22.5V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H6.455Zm-.692-2H20V5H4v13.385L5.763 17ZM7 10h2a3 3 0 1 0 6 0h2a5 5 0 0 1-10 0Z"/>`),
		g.Group(children),
	)
}