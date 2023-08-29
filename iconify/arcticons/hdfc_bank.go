package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdfcBank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30 18H18v12h12Zm12.5 3V7.5a2 2 0 0 0-2-2H27m-6 0H7.5a2 2 0 0 0-2 2V21m0 6v13.5a2 2 0 0 0 2 2H21m6 0h13.5a2 2 0 0 0 2-2V27"/>`),
		g.Group(children),
	)
}