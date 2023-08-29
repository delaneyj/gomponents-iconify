package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BalloonOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 8a2 2 0 0 0-2-2"/><path d="M7.762 3.753A6 6 0 0 1 18 8c0 1.847-.37 3.564-1.007 4.993m-1.59 2.42C14.436 16.413 13.263 17 12 17c-3.314 0-6-4.03-6-9c0-.593.086-1.166.246-1.707M12 17v1a2 2 0 0 1-2 2H7a2 2 0 0 0-2 2M3 3l18 18"/></g>`),
		g.Group(children),
	)
}