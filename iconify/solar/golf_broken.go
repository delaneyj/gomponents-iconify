package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GolfBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M16.979 21.536C15.513 21.831 13.813 22 12 22c-5.523 0-10-1.567-10-3.5S6.477 15 12 15s10 1.567 10 3.5c0 .548-.36 1.066-1 1.527M12 18V2m0 1.5l5.422 2.711c1.561.78 2.342 1.171 2.342 1.789c0 .618-.78 1.008-2.342 1.789L12 12.5"/>`),
		g.Group(children),
	)
}