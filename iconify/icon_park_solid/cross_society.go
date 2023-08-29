package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossSociety(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCrossSociety0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M7 18h11V7a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v11h11a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H30v11a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2V30H7a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCrossSociety0)"/>`),
		g.Group(children),
	)
}