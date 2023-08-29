package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresentationLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 14h-1V3a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v11H3a1 1 0 0 0 0 2h8v1.15l-4.55 3A1 1 0 0 0 7 22a.94.94 0 0 0 .55-.17L11 19.55V21a1 1 0 0 0 2 0v-1.45l3.45 2.28A.94.94 0 0 0 17 22a1 1 0 0 0 .55-1.83l-4.55-3V16h8a1 1 0 0 0 0-2Zm-3 0H6V4h12Zm-9-2a1 1 0 0 0 .83-.45l1.33-2l1.13 1.14a1 1 0 0 0 .81.29a1 1 0 0 0 .73-.45l2-3a1 1 0 0 0-1.66-1.1l-1.33 2l-1.13-1.14A1 1 0 0 0 10.9 7a1 1 0 0 0-.73.45l-2 3a1 1 0 0 0 .28 1.38A.94.94 0 0 0 9 12Z"/>`),
		g.Group(children),
	)
}