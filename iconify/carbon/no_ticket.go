package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoTicket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 6h-.586L30 3.414L28.586 2L2 28.586L3.414 30l4-4H28a2.003 2.003 0 0 0 2-2v-5a1 1 0 0 0-1-1a2 2 0 0 1 0-4a1 1 0 0 0 1-1V8a2.002 2.002 0 0 0-2-2zm0 6.127a4 4 0 0 0 0 7.746V24h-7v-3h-2v3H9.414L19 14.414V19h2v-6.586L25.414 8H28zm-24 0V8h15V6H4a2.002 2.002 0 0 0-2 2v5a1 1 0 0 0 1 1a2 2 0 0 1 0 4a1 1 0 0 0-1 1v5h2v-4.127a4 4 0 0 0 0-7.747z"/>`),
		g.Group(children),
	)
}