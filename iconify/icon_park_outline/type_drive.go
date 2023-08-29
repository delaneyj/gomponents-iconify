package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TypeDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 12a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v24a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V12Z"/><path d="M15.414 32.586A2 2 0 0 1 16.828 32H30.26a2 2 0 0 1 1.302.481L38 38H10l5.414-5.414ZM39 38H9"/><circle cx="15" cy="21" r="4"/><circle cx="33" cy="21" r="4"/><path d="M15 25h18m-18-8h18"/></g>`),
		g.Group(children),
	)
}