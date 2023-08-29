package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoursoramaBanque(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.584 41.5a6 6 0 0 1-6-6v-17h-17a6 6 0 0 1 0-12h23a6 6 0 0 1 6 6v23a6 6 0 0 1-6 6Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.584 41.5a6 6 0 0 1-4.242-10.242l23-23a6 6 0 1 1 8.484 8.484l-23 23a5.979 5.979 0 0 1-4.242 1.758Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.584 18.5v-6a6 6 0 0 1 6-6h0m-6 12h6a6 6 0 0 0 6-6h0"/>`),
		g.Group(children),
	)
}