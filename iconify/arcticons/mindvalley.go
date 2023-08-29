package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mindvalley(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.5 32.555v-8.392a5.853 5.853 0 0 0-3.76-5.473c-3.116-1.193-8.035-2.679-14.24-3.246c2.828 6.165 8.259 12.648 18 17.111Zm3 0v-8.392a5.853 5.853 0 0 1 3.76-5.473c3.116-1.193 8.035-2.679 14.24-3.246c-2.828 6.165-8.259 12.648-18 17.111Z"/>`),
		g.Group(children),
	)
}