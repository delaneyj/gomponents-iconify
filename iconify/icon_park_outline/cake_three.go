package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CakeThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M39 34V16a2 2 0 0 0-2-2H11a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2h26a2 2 0 0 0 2-2ZM6 36h36M6 42h36"/><path d="M9 21h6v5h6v-5h6v8h6v-8h6M9 23v-4m30 4v-4m-6-5V8m-9 6V4m-9 10V8"/></g>`),
		g.Group(children),
	)
}