package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gramly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.138 24.369l3.142-7.434l-9.258 21.904c-2.553 6.04-7.15 5.384-8.245 2.664c-.558-1.385-1.443-4.794 6.472-6.144c6.33-1.08 14.355-3.457 18.32-12.837m-10.43 1.847c-1.586 3.75-5.925 6.792-9.693 6.792s-5.537-3.041-3.951-6.792m19.17-13.078c1.586-3.75-.183-6.791-3.95-6.791s-8.107 3.04-9.693 6.792l-5.527 13.077"/>`),
		g.Group(children),
	)
}