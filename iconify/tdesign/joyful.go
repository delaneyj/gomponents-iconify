package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Joyful(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm6.619-4.862L11 9.04v1.46l-3.2 2.4l-1.2-1.6l1.86-1.394L6.637 8.88l.98-1.743Zm9.743 1.743L15.54 9.906L17.4 11.3l-1.2 1.6l-3.2-2.4V9.04l3.381-1.902l.98 1.743ZM7 13h10v1a5 5 0 0 1-10 0v-1Zm2.17 2a3.001 3.001 0 0 0 5.66 0H9.17Z"/>`),
		g.Group(children),
	)
}