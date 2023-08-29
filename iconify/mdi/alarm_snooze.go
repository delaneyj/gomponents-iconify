package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmSnooze(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.88 3.39L6.6 1.86L2 5.71l1.29 1.53l4.59-3.85M22 5.72l-4.6-3.86l-1.29 1.53l4.6 3.86L22 5.72M12 4a9 9 0 0 0-9 9a9 9 0 0 0 9 9a9 9 0 0 0 9-9a9 9 0 0 0-9-9m0 16a7 7 0 0 1-7-7a7 7 0 0 1 7-7a7 7 0 0 1 7 7a7 7 0 0 1-7 7m-3-9h3.63L9 15.2V17h6v-2h-3.63L15 10.8V9H9v2Z"/>`),
		g.Group(children),
	)
}