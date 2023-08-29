package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RobotOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 7h6a2 2 0 0 1 2 2v1l1 1v3l-1 1m-.171 3.811A2 2 0 0 1 17 20H7a2 2 0 0 1-2-2v-3l-1-1v-3l1-1V9a2 2 0 0 1 2-2m3 9h4"/><path d="M7.5 11.5a1 1 0 1 0 2 0a1 1 0 1 0-2 0m8.354.353A.498.498 0 0 0 15.5 11a.498.498 0 0 0-.356.149M8.336 4.343L8 3m7 4l1-4M3 3l18 18"/></g>`),
		g.Group(children),
	)
}