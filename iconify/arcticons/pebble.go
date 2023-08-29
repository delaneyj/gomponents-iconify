package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pebble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41 5H7a2 2 0 0 0-2 2v34a2 2 0 0 0 2 2h34a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 9.5A14.5 14.5 0 1 0 38.5 24A14.5 14.5 0 0 0 24 9.5ZM24 24v-8.47m7.68 14.36L24 24"/>`),
		g.Group(children),
	)
}