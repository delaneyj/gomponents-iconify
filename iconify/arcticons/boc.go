package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 24A21.5 21.5 0 1 1 24 2.5A21.47 21.47 0 0 1 45.5 24Z"/><path fill="none" stroke="currentColor" d="M19.8 20h8.4a1.16 1.16 0 0 1 1.2 1.2v5.6a1.16 1.16 0 0 1-1.2 1.2h-8.4a1.16 1.16 0 0 1-1.2-1.2v-5.6a1.16 1.16 0 0 1 1.2-1.2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.1 24a8.42 8.42 0 0 1 7.2-8.4V7.8a16.43 16.43 0 0 0 0 32.4v-7.8a8.42 8.42 0 0 1-7.2-8.4Zm19.8 0a8.42 8.42 0 0 1-7.2 8.4v7.8a16.43 16.43 0 0 0 0-32.4v7.8a8.42 8.42 0 0 1 7.2 8.4Z"/>`),
		g.Group(children),
	)
}