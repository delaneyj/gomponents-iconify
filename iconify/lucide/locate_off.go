package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocateOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12h3m14 0h3M12 2v3m0 14v3M7.11 7.11C5.83 8.39 5 10.1 5 12c0 3.87 3.13 7 7 7c1.9 0 3.61-.83 4.89-2.11m1.82-2.93c.19-.63.29-1.29.29-1.96c0-3.87-3.13-7-7-7c-.67 0-1.33.1-1.96.29M2 2l20 20"/>`),
		g.Group(children),
	)
}