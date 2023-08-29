package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BbcSport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 6.078h37v11.948h-37zm13.029 11.948H42.5v11.948H18.529zm10.603 11.948H42.5v11.948H29.132z"/>`),
		g.Group(children),
	)
}