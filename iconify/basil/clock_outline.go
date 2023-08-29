package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.75 7a.75.75 0 0 0-1.5 0v5a.75.75 0 0 0 .352.636l3 1.875a.75.75 0 1 0 .796-1.272l-2.648-1.655V7Z"/><path fill="currentColor" fill-rule="evenodd" d="M12 3.25a8.75 8.75 0 1 0 0 17.5a8.75 8.75 0 0 0 0-17.5ZM4.75 12a7.25 7.25 0 1 1 14.5 0a7.25 7.25 0 0 1-14.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}