package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8.618V17.5m0-8.882a1 1 0 0 0-.553-.894l-.491-.246a7 7 0 0 0-4.12-.669l-.773.11A8 8 0 0 1 4.93 7H4v10h.38a8 8 0 0 0 2.197-.308l.553-.158a5 5 0 0 1 3.61.336l1.26.63m0-8.882a1 1 0 0 1 .553-.894l.491-.246a7 7 0 0 1 4.12-.669l.773.11c.375.054.753.081 1.131.081H20v10h-.38a8 8 0 0 1-2.197-.308l-.553-.158a5 5 0 0 0-3.61.336L12 17.5"/>`),
		g.Group(children),
	)
}