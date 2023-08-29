package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dailydozen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.743 43.5h0a1.969 1.969 0 0 0 1.839-1.266L27.277 6.399A2.953 2.953 0 0 1 30.035 4.5h9.943a1.969 1.969 0 0 0-1.839 1.266L24.444 41.601a2.953 2.953 0 0 1-2.758 1.899h-9.943"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.956 33.404L14.97 14.17a1 1 0 0 0-1.895-.332L8.39 23.526a4 4 0 0 0-.388 2.037l1.187 16.012a2 2 0 0 0 1.999 1.925h.555"/>`),
		g.Group(children),
	)
}