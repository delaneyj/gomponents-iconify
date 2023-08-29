package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandleRound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHandleRound0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20" fill="#555"/><circle cx="24" cy="24" r="12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHandleRound0)"/>`),
		g.Group(children),
	)
}