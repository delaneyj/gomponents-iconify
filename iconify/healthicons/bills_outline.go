package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BillsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M10.49 32v4a2 2 0 0 0 2 2h30a2 2 0 0 0 2-2V20a2 2 0 0 0-2-2H40v3.669a2.994 2.994 0 0 0 2.49 1.327V33a3 3 0 0 0-2.983 3H15.485a3.002 3.002 0 0 0-2.996-2.962V32h-2ZM21 22a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 2a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M6 10a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h30a2 2 0 0 0 2-2V12a2 2 0 0 0-2-2H6Zm2.998 2h24.003A2.999 2.999 0 0 0 36 14.996V25a2.998 2.998 0 0 0-2.983 3H8.996A3 3 0 0 0 6 25.038V15a2.998 2.998 0 0 0 2.998-3Z"/></g>`),
		g.Group(children),
	)
}