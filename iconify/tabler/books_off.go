package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BooksOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 9v10a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V5m3-1a1 1 0 0 1 1 1m0 0a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v4m0 4v6a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1V9M5 8h3m1 8h4"/><path d="M14.254 10.244L13.036 5.82a1.02 1.02 0 0 1 .634-1.219l.133-.041l2.184-.53c.562-.135 1.133.19 1.282.732l3.236 11.75m-.92 3.077l-1.572.38c-.562.136-1.133-.19-1.282-.731l-.952-3.458M14 9l4-1m1.207 7.199l.716-.18M3 3l18 18"/></g>`),
		g.Group(children),
	)
}