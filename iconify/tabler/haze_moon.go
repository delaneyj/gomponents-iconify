package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HazeMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 16h18M3 20h18M8.296 16c-2.268-1.4-3.598-4.087-3.237-6.916c.443-3.48 3.308-6.083 6.698-6.084v.006h.296c-1.991 1.916-2.377 5.03-.918 7.405c1.459 2.374 4.346 3.33 6.865 2.275A6.888 6.888 0 0 1 15.223 16"/>`),
		g.Group(children),
	)
}