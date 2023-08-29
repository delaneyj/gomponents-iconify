package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mercari(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.154 43.5V14.991L22.408 43.5l14.255-28.509V43.5M33.3 9.663L34.684 4.5l5.164 1.383l-1.384 5.164z"/>`),
		g.Group(children),
	)
}