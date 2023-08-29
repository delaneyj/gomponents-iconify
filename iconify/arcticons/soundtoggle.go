package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Soundtoggle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.557 42.5L15.724 28.666H6.39v-9.332h9.332L29.557 5.5Zm3.85-8.591a10.086 10.086 0 0 0 0-19.818V33.91Z"/>`),
		g.Group(children),
	)
}