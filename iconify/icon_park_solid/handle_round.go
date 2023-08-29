package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandleRound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHandleRound0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20" fill="#fff" stroke="#fff"/><circle cx="24" cy="24" r="12" stroke="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHandleRound0)"/>`),
		g.Group(children),
	)
}