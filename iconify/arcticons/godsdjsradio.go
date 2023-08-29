package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Godsdjsradio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.68 41.673V23.091h-7.432v-4.646h7.432v-6.504h4.645v6.504h7.432v4.646h-7.432v18.582Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.972 15.793S10.58 6.2 24.089 6.328s12.94 9.744 12.94 9.744m-.027 9.884a4.817 4.817 0 0 0 4.698-4.928a4.82 4.82 0 0 0-4.671-4.956Zm-26.004-.281A4.817 4.817 0 0 1 6.3 20.747a4.82 4.82 0 0 1 4.671-4.955ZM6.3 20.747h-.8m37 0h-.8"/>`),
		g.Group(children),
	)
}