package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileExcel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feFileExcel0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFileExcel1" fill="currentColor"><path id="feFileExcel2" d="M6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm9 2H6v16h12V7h-3V4Zm-8 8h3l2 2l2-2h3l-3 3l3 3h-3l-2-2l-2 2H7l3-3l-3-3Z"/></g></g>`),
		g.Group(children),
	)
}