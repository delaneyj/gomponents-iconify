package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockingComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 6H9a3 3 0 0 0-3 3v22a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3v-4m-18 7v8"/><rect width="12" height="8" x="30" y="12" rx="3"/><path d="M36 6a3 3 0 0 1 3 3v3h-6V9a3 3 0 0 1 3-3Z"/><path stroke-linecap="round" d="M14 42h20"/></g>`),
		g.Group(children),
	)
}