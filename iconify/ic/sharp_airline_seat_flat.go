package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpAirlineSeatFlat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 7v6H9V7h13zM2 14v2h6v2h8v-2h6v-2H2zm5.14-1.9a3 3 0 0 0-.04-4.24a3 3 0 0 0-4.24.04a3 3 0 0 0 .04 4.24a3 3 0 0 0 4.24-.04z"/>`),
		g.Group(children),
	)
}