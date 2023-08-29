package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KavliIpmu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2h0v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM11.474 26.956c0 .715-.579 1.294-1.293 1.294h0a1.293 1.293 0 0 1-1.293-1.294h0a1.293 1.293 0 0 1 2.586 0Zm19.665 3.438v5.3c0 1.5 1.2 2.7 2.6 2.7s2.6-1.2 2.6-2.7v-5.3m-15.353 8v-8l4 8l4-8v8m-15.452 0v-8h2.6c1.5 0 2.7 1.2 2.7 2.7s-1.2 2.7-2.7 2.7h-2.6m-3.353-5.4v8m-1.2-8h2.4m-2.4 8h2.4"/><path fill="currentColor" d="M16.87 9.974a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm21.164 15.403a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm-9.329-2.925a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Z"/><circle cx="26.953" cy="19.585" r=".75" fill="currentColor"/><path fill="currentColor" d="M37.127 11.346a.75.75 0 0 1-1.5 0a.75.75 0 0 1 1.5 0Zm-4.035-.462a.75.75 0 0 1-1.5 0a.75.75 0 0 1 1.5 0Z"/>`),
		g.Group(children),
	)
}