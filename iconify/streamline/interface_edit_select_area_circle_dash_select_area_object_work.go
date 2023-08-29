package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditSelectAreaCircleDashSelectAreaObjectWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5.67a6.7 6.7 0 0 0-3 0m-1.91.8a7 7 0 0 0-1.19.93a7 7 0 0 0-.93 1.19M.67 5.5a6.7 6.7 0 0 0 0 3m.8 1.91a7 7 0 0 0 .93 1.19a7 7 0 0 0 1.19.93m1.91.8a6.7 6.7 0 0 0 3 0m1.91-.8a7 7 0 0 0 1.19-.93a7 7 0 0 0 .93-1.19m.8-1.91a6.7 6.7 0 0 0 0-3m-.8-1.91a7 7 0 0 0-.93-1.19a7 7 0 0 0-1.19-.93"/>`),
		g.Group(children),
	)
}