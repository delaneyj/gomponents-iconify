package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaptureFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M8 3a1 1 0 0 1 .117 1.993L8 5H6a1 1 0 0 0-.993.883L5 6v2a1 1 0 0 1-1.993.117L3 8V6a3 3 0 0 1 2.824-2.995L6 3h2zM4 15a1 1 0 0 1 .993.883L5 16v2a1 1 0 0 0 .883.993L6 19h2a1 1 0 0 1 .117 1.993L8 21H6a3 3 0 0 1-2.995-2.824L3 18v-2a1 1 0 0 1 1-1zM18 3a3 3 0 0 1 2.995 2.824L21 6v2a1 1 0 0 1-1.993.117L19 8V6a1 1 0 0 0-.883-.993L18 5h-2a1 1 0 0 1-.117-1.993L16 3h2zm2 12a1 1 0 0 1 .993.883L21 16v2a3 3 0 0 1-2.824 2.995L18 21h-2a1 1 0 0 1-.117-1.993L16 19h2a1 1 0 0 0 .993-.883L19 18v-2a1 1 0 0 1 1-1zm-8-7a4 4 0 1 1-3.995 4.2L8 12l.005-.2A4 4 0 0 1 12 8z"/></g>`),
		g.Group(children),
	)
}