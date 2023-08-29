package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nextcloudsms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.55 5.43a2.05 2.05 0 0 0-2 2.05v22.58a2.05 2.05 0 0 0 2.05 2.06h6l11.59 8a1 1 0 0 0 1.57-.82v-7.18h15.69a2.05 2.05 0 0 0 2-2.06V7.48a2.05 2.05 0 0 0-2-2.05Zm1.08 6.13h32.74M7.63 18.77h32.74M7.63 25.99H20.9"/>`),
		g.Group(children),
	)
}