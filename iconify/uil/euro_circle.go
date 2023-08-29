package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuroCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1a11 11 0 1 0 11 11A11 11 0 0 0 12 1Zm0 20a9 9 0 1 1 9-9a9 9 0 0 1-9 9Zm.59-13.33a3.34 3.34 0 0 1 2.62 1.38a1 1 0 0 0 1.4.19a1 1 0 0 0 .18-1.41a5.32 5.32 0 0 0-4.2-2.16A5.57 5.57 0 0 0 7.46 9.5H6a1 1 0 0 0 0 2h1v1H6a1 1 0 0 0 0 2h1.46a5.57 5.57 0 0 0 5.13 3.83a5.32 5.32 0 0 0 4.2-2.16A1 1 0 1 0 15.21 15a3.34 3.34 0 0 1-2.62 1.38a3.42 3.42 0 0 1-2.92-1.88H12a1 1 0 0 0 0-2H9.05A4.23 4.23 0 0 1 9 12a4.23 4.23 0 0 1 .05-.5H12a1 1 0 0 0 0-2H9.67a3.42 3.42 0 0 1 2.92-1.83Z"/>`),
		g.Group(children),
	)
}