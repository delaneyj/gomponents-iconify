package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterPercent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3.25S6 10 6 14c0 3.32 2.69 6 6 6a6 6 0 0 0 6-6c0-4-6-10.75-6-10.75m2.47 6.72l1.06 1.06l-6 6l-1.06-1.06M9.75 10A1.25 1.25 0 0 1 11 11.25a1.25 1.25 0 0 1-1.25 1.25a1.25 1.25 0 0 1-1.25-1.25A1.25 1.25 0 0 1 9.75 10m4.5 4.5a1.25 1.25 0 0 1 1.25 1.25A1.25 1.25 0 0 1 14.25 17A1.25 1.25 0 0 1 13 15.75a1.25 1.25 0 0 1 1.25-1.25Z"/>`),
		g.Group(children),
	)
}