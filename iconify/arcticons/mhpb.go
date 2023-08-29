package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mhpb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.903 36.46c16.416 0 21.69-23.475 21.69-30.96h12.503c0 12.759-8.76 37-34.193 37v-6.04Z"/>`),
		g.Group(children),
	)
}