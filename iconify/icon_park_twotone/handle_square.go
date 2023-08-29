package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandleSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHandleSquare0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20" fill="#555"/><path d="M14 14h20v20H14z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHandleSquare0)"/>`),
		g.Group(children),
	)
}