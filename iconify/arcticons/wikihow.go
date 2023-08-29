package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wikihow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2ZM12 26.1V36m6.5-9.9V36M12 31h6.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.8 15.9l-2 6.6l-2.1-6.6l-2 6.6l-2.1-6.6M23.8 36h0c-1.4 0-2.5-1.1-2.5-2.5v-1.6c0-1.4 1.1-2.5 2.5-2.5h0c1.4 0 2.5 1.1 2.5 2.5v1.6c0 1.4-1.1 2.5-2.5 2.5Zm3-23.4v9.9m0-2.1l4.5-4.5m-3 3.1l3.5 3.5"/><circle cx="23" cy="12.9" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23 15.9v6.6"/><circle cx="34.9" cy="12.9" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.9 15.9v6.6m1.5 6.9l-2 6.6l-2.1-6.6l-2 6.6l-2-6.6"/>`),
		g.Group(children),
	)
}