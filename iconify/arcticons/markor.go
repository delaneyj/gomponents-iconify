package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Markor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 5a2 2 0 0 0-2 2v6.58h38V7a2 2 0 0 0-2-2Zm-2 8.58V41a2 2 0 0 0 2 2h34a2 2 0 0 0 2-2V13.58Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.98 37.26V21.18L24 37.28l8.02-16.07v16.07M10.93 7.67v3.24m6.53-3.24v3.24M24 7.67v3.24m6.54-3.24v3.24m6.53-3.24v3.24"/>`),
		g.Group(children),
	)
}