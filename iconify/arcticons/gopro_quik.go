package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoproQuik(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.244 22.346v5.077c0 1.7 1.378 3.077 3.077 3.077h0c1.7 0 3.077-1.378 3.077-3.077v-5.077m0 5.077V30.5m6.365-12.308V30.5m.001-2.614l5.571-5.544m-3.797 3.779l4.379 4.361"/><circle cx="29.471" cy="18.577" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.471 22.346V30.5"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m17.237 30.5l-4.077-4.028"/><rect width="8.154" height="12.308" x="9.083" y="18.192" rx="4.077" ry="4.077"/></g>`),
		g.Group(children),
	)
}