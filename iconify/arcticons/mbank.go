package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mbank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.71 22.27a5.75 5.75 0 0 1 5.76-5.75h0a5.74 5.74 0 0 1 5.75 5.75v9.21"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.71 31.48V16.52H9.5m14.72 5.75A5.75 5.75 0 0 1 30 16.52h0a5.74 5.74 0 0 1 5.75 5.75v7.33a1.88 1.88 0 0 0 2.77 1.65h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}