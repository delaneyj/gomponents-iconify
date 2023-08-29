package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerAuto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 14c-.3-.4-.6-.7-1-1V5c0-1.7-1.3-3-3-3S5 3.3 5 5v8c-2.2 1.7-2.7 4.8-1 7s4.8 2.7 7 1s2.7-4.8 1-7M9 8H7V5c0-.5.5-1 1-1s1 .5 1 1v3m9-5h-2l-3.2 9h1.9l.7-2h3.2l.7 2h1.9L18 3m-2.2 5.7L17 5l1.2 3.7h-2.4Z"/>`),
		g.Group(children),
	)
}