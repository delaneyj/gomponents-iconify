package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" fill-rule="evenodd" d="M24 16H12V32C12 38.6274 17.3726 44 24 44C30.6274 44 36 38.6274 36 32V16H24Z" clip-rule="evenodd"/><path d="M36 16C36 9.37258 30.6274 4 24 4V16H36Z"/><path d="M24 4C17.3726 4 12 9.37258 12 16H24V4Z"/></g>`),
		g.Group(children),
	)
}