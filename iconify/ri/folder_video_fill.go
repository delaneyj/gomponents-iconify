package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderVideoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 5h-8.586l-2-2H3a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1Zm-6 7.667a.4.4 0 0 1 0 .666l-4.878 3.252a.4.4 0 0 1-.622-.333V9.747a.4.4 0 0 1 .622-.333L15 12.667Z"/>`),
		g.Group(children),
	)
}