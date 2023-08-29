package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18Zm11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11Zm-9.268 5a2 2 0 0 1-3.464 0H6v-2.736l1-2V10a5 5 0 0 1 10 0v2.264l1 2V17h-4.268ZM8 15h8v-.264l-1-2V10a3 3 0 1 0-6 0v2.736l-1 2V15Z"/>`),
		g.Group(children),
	)
}