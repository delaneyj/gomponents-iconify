package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.53 3.47a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 0 1-1.06-1.06l2.5-2.5a.75.75 0 0 1 1.06 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M12 4.5a8.5 8.5 0 1 0 0 17a8.5 8.5 0 0 0 0-17Zm.75 3.5a.75.75 0 0 0-1.5 0v5a.75.75 0 0 0 .352.636l3 1.875a.75.75 0 1 0 .796-1.272l-2.648-1.655V8Z" clip-rule="evenodd"/><path fill="currentColor" d="M17.47 4.53a.75.75 0 0 1 1.06-1.06l2.5 2.5a.75.75 0 0 1-1.06 1.06l-2.5-2.5Z"/>`),
		g.Group(children),
	)
}