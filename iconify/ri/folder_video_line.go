package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderVideoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19V5h5.586l2 2H20v12H4ZM21 5h-8.586l-2-2H3a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1Zm-6 7.667l-4.878-3.253a.4.4 0 0 0-.622.333v6.505a.4.4 0 0 0 .622.333L15 13.333a.401.401 0 0 0 0-.666Z"/>`),
		g.Group(children),
	)
}