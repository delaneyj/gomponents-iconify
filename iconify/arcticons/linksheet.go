package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Linksheet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.253 23.954l14.475.037m-11.096 8.69h-6.96a8.681 8.681 0 1 1 0-17.362h6.644m8.034 0h6.96a8.681 8.681 0 1 1 0 17.362h-6.645"/>`),
		g.Group(children),
	)
}