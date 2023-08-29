package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Atk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM19.222 17.871h8.107m-4.054 12.237V17.871m-6.011 8.184h-5.421"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.5 30.108l4.053-12.237l4.054 12.237m12.316-12.216v12.237m0-4.262l6.577-7.934m0 12.196l-5.038-6.119"/>`),
		g.Group(children),
	)
}