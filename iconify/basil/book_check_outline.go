package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookCheckOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.541 7.48a.75.75 0 0 1-.022 1.061l-3.125 3a.75.75 0 0 1-1.038 0l-1.875-1.8a.75.75 0 1 1 1.038-1.082l1.356 1.301l2.606-2.501a.75.75 0 0 1 1.06.022Z"/><path fill="currentColor" fill-rule="evenodd" d="M3.75 8A4.75 4.75 0 0 1 8.5 3.25h10c.966 0 1.75.784 1.75 1.75v15a1.75 1.75 0 0 1-1.75 1.75h-11A3.75 3.75 0 0 1 3.75 18V8Zm1.5 7a3.734 3.734 0 0 1 2.25-.75h11.25V5a.25.25 0 0 0-.25-.25h-10A3.25 3.25 0 0 0 5.25 8v7Zm0 3a2.25 2.25 0 0 0 2.25 2.25h11a.25.25 0 0 0 .25-.25v-4.25H7.5A2.25 2.25 0 0 0 5.25 18Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}