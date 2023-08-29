package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xdadevelopers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.54 16.61v14.78"/><rect width="7.39" height="9.83" x="20.15" y="21.56" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.7"/><rect width="7.39" height="9.83" x="31.11" y="21.56" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.7"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.5 21.56v9.83m-21.61-9.83L9.5 31.39m7.39 0L9.5 21.56m31-16.06h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}