package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M7 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0-3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm3-3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm3-3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path stroke="currentColor" stroke-width="1.5" d="M2 11c0-2.828 0-4.243.879-5.121C3.757 5 5.172 5 8 5h8c2.828 0 4.243 0 5.121.879C22 6.757 22 8.172 22 11v2c0 2.828 0 4.243-.879 5.121C20.243 19 18.828 19 16 19H8c-2.828 0-4.243 0-5.121-.879C2 17.243 2 15.828 2 13v-2Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M7 16h10"/></g>`),
		g.Group(children),
	)
}