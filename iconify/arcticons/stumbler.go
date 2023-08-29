package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stumbler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.19 12C9.87 7.35 16.65 4.5 24 4.5S38.13 7.35 41.81 12m-27.92 7c2.11-2.48 5.9-4 10.05-4s8 1.5 10.05 3.95m-10.05 4.51a7.38 7.38 0 0 0-6.06 11.6l6.06 8.44l6-8.44h0a7.38 7.38 0 0 0-6-11.6Zm0 4.92a2.46 2.46 0 0 1 2.46 2.46h0a2.46 2.46 0 0 1-2.46 2.46h0a2.47 2.47 0 0 1-2.47-2.46h0a2.47 2.47 0 0 1 2.47-2.46Z"/>`),
		g.Group(children),
	)
}