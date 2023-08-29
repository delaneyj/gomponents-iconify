package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 21H9V10h6v11Zm2 0V10h5v10a1 1 0 0 1-1 1h-4ZM7 21H3a1 1 0 0 1-1-1V10h5v11ZM22 8H2V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v4Z"/>`),
		g.Group(children),
	)
}