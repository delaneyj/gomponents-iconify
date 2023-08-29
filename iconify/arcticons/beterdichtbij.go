package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beterdichtbij(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.964 5.5L24.055 17.15L12.066 5.5M42.5 11.667l-11.428 11.61L42.5 35.288M36.084 42.5L24.055 30.77L11.825 42.5M5.5 11.266l11.708 12.01L5.5 35.45"/>`),
		g.Group(children),
	)
}