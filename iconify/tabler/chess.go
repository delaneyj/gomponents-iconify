package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3a3 3 0 0 1 3 3c0 1.113-.6 2.482-1.5 3l1.5 7H9l1.5-7C9.6 8.482 9 7.113 9 6a3 3 0 0 1 3-3zM8 9h8m-9.316 7.772a1 1 0 0 0-.684.949V19a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-1.28a1 1 0 0 0-.684-.948L15 16H9l-2.316.772z"/>`),
		g.Group(children),
	)
}