package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 3a1 1 0 0 1 1-1h.01a1 1 0 0 1 0 2H7a1 1 0 0 1-1-1Zm2 3a1 1 0 0 0-2 0v1a2 2 0 0 0-2 2v1a2 2 0 0 0-2 2v.683a3.7 3.7 0 0 1 1.055.485a1.704 1.704 0 0 0 1.89 0a3.704 3.704 0 0 1 4.11 0a1.704 1.704 0 0 0 1.89 0a3.704 3.704 0 0 1 4.11 0a1.704 1.704 0 0 0 1.89 0A3.7 3.7 0 0 1 18 12.683V12a2 2 0 0 0-2-2V9a2 2 0 0 0-2-2V6a1 1 0 1 0-2 0v1h-1V6a1 1 0 1 0-2 0v1H8V6Zm10 8.868a3.704 3.704 0 0 1-4.055-.036a1.704 1.704 0 0 0-1.89 0a3.704 3.704 0 0 1-4.11 0a1.704 1.704 0 0 0-1.89 0A3.704 3.704 0 0 1 2 14.868V17a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-2.132ZM9 3a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H10a1 1 0 0 1-1-1Zm3 0a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H13a1 1 0 0 1-1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}