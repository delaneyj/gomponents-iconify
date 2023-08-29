package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleCardboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.202 12.593v-.919a2.16 2.16 0 0 0-2.16-2.16H9.958a2.16 2.16 0 0 0-2.16 2.16v.919m33.542 0H6.66a2.16 2.16 0 0 0-2.16 2.16v21.573a2.16 2.16 0 0 0 2.16 2.16h9.432c3.584 0 3.483-9.868 7.908-9.868s4.324 9.868 7.908 9.868h9.432a2.16 2.16 0 0 0 2.16-2.16V14.753a2.16 2.16 0 0 0-2.16-2.16Z"/><circle cx="14.774" cy="25.022" r="4.726" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.226" cy="25.022" r="4.726" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.165 34.978h9.67"/>`),
		g.Group(children),
	)
}