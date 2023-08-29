package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Round(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRound0"><circle cx="24" cy="24" r="20" fill="#555" stroke="#fff" stroke-width="4"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRound0)"/>`),
		g.Group(children),
	)
}