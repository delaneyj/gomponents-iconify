package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Greenlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.29 13.79v13.89a6.529 6.529 0 0 0 6.53 6.53h3.68M24.16 20.635a6.837 6.837 0 0 0-7.197-6.835a7.1 7.1 0 0 0-6.463 7.222v6.343a6.837 6.837 0 0 0 6.83 6.845h0a6.837 6.837 0 0 0 6.83-6.845h-6.83"/>`),
		g.Group(children),
	)
}