package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDownloadAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 15.5h-4c-1.152-.419-2-1.703-2-3a3 3 0 0 1 3.1-2.998a5.002 5.002 0 1 1 9.702 2.397A2 2 0 1 1 16.5 15.5h-4m-4 2l2 2l2-2m-2-8v10"/>`),
		g.Group(children),
	)
}