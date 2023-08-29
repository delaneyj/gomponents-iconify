package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Novaprime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.802 18.694l-7.838-5.898l-4.239 1.135V9.544L24.001 3.5l-7.726 6.044v4.387l-4.237-1.135l-7.839 5.898L6.455 24L4.2 29.307l7.839 5.896l4.237-1.136v4.388l7.726 6.045l7.724-6.045v-4.388l4.24 1.136l7.837-5.898L41.545 24ZM20.11 22.67v6.549h7.778V22.67m1.611 1.611l-5.5-5.5l-5.5 5.5m-2.224-10.35v20.136m15.45-20.136v20.136"/>`),
		g.Group(children),
	)
}