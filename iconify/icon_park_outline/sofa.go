package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sofa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M12 21H4v14h8V21Zm32 0h-8v14h8V21Z"/><path stroke-linecap="round" d="M36 27H12v8h24v-8ZM8 20V8h32v12M8 36v4m32-4v4"/></g>`),
		g.Group(children),
	)
}