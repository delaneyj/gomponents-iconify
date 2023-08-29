package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSum0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 7v12H7a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V29h12a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H21a2 2 0 0 0-2 2Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSum0)"/>`),
		g.Group(children),
	)
}