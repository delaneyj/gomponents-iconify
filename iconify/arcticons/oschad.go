package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oschad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.345 5.5C30.301 12.736 19.363 31.899 5.5 22.805"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.1 5.5c7.904 18.65-9.691 34.218-26.6 28.502"/>`),
		g.Group(children),
	)
}