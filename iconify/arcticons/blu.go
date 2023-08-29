package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.46 19.712v5.808a3.52 3.52 0 0 0 3.52 3.52h0a3.52 3.52 0 0 0 3.52-3.52v-5.808m0 5.808v3.521M23.438 14.96v12.32c0 .973.788 1.76 1.76 1.76h.528M12.5 23.232a3.52 3.52 0 0 1 3.52-3.52h0a3.52 3.52 0 0 1 3.52 3.52v2.288a3.52 3.52 0 0 1-3.52 3.52h0a3.52 3.52 0 0 1-3.52-3.52m0 3.521V14.959"/>`),
		g.Group(children),
	)
}