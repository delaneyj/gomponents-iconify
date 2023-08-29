package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserLove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M12 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M22 17.28a2.28 2.28 0 0 1-.662 1.606c-.976.984-1.923 2.01-2.936 2.958a.597.597 0 0 1-.823-.017l-2.918-2.94a2.281 2.281 0 0 1 0-3.214a2.277 2.277 0 0 1 3.233 0l.106.107l.106-.107A2.277 2.277 0 0 1 22 17.28Z"/><path stroke-linecap="round" d="M5 20v-1a7 7 0 0 1 10-6.326"/></g>`),
		g.Group(children),
	)
}