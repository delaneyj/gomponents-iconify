package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataRefinery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7 30a5 5 0 1 1 5-5a5.006 5.006 0 0 1-5 5Zm0-8a3 3 0 1 0 3 3a3.003 3.003 0 0 0-3-3Z"/><path fill="currentColor" d="M24 26H14v-2h10v-7H8a2.002 2.002 0 0 1-2-2V8a2.002 2.002 0 0 1 2-2h10v2H8v7h16a2.002 2.002 0 0 1 2 2v7a2.002 2.002 0 0 1-2 2Z"/><path fill="currentColor" d="M28 12h-6a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h6a2.002 2.002 0 0 1 2 2v6a2.002 2.002 0 0 1-2 2Zm-6-8v6h6V4Z"/>`),
		g.Group(children),
	)
}