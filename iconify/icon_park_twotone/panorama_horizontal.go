package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPanoramaHorizontal0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 11s9 4 20 4s20-4 20-4v26s-9-4-20-4s-20 4-20 4V11Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPanoramaHorizontal0)"/>`),
		g.Group(children),
	)
}