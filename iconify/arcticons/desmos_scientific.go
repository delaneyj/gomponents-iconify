package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesmosScientific(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.618 18.847a1.701 1.701 0 0 1 1.25-.35h.4a1.003 1.003 0 0 1 1 1h0a1.003 1.003 0 0 1-1 1h0a1 1 0 0 1 0 2h-.4a1.627 1.627 0 0 1-1.25-.35m15.582 7.35l4.3-8h-5.3m-7.232 2.6a2.67 2.67 0 0 1 2.6-2.7a2.687 2.687 0 0 1 1.9 4.6l-4.5 3.5h5.3m-13.041-3l.755-2h.199l1.887 5h.201l3.774-10H33.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.504h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}