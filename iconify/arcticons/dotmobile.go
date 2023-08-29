package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dotmobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.157 35.113a12.285 12.285 0 1 1 11.89-.112"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.153 42.56a21.498 21.498 0 1 1 22.009-.188"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.784 28.63a4.915 4.915 0 1 1 .481.228"/>`),
		g.Group(children),
	)
}