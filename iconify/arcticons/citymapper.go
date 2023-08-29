package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Citymapper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.62 20.84A3.12 3.12 0 1 1 4.5 24a3.12 3.12 0 0 1 3.12-3.12Zm32.76 0A3.12 3.12 0 1 1 37.26 24a3.12 3.12 0 0 1 3.12-3.12Zm-26.74 3.14h20.72m-6.31-6.3l6.31 6.3m-6.31 6.3l6.31-6.3"/>`),
		g.Group(children),
	)
}