package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monespace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m3.5 31.422l41-9.778H19.523s-.942 4.359-16.023 9.778Zm29.135-6.948a51.242 51.242 0 0 1-6.146 6.948H3.5m16.023-9.778c-1.178-1.768-4.713-3.888-8.6-5.066h23.68a2.47 2.47 0 0 1 .81 2.061a6.954 6.954 0 0 1-.976 3.005"/>`),
		g.Group(children),
	)
}