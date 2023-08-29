package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandleSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHandleSquare0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20" fill="#fff" stroke="#fff"/><path stroke="#000" d="M14 14h20v20H14z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHandleSquare0)"/>`),
		g.Group(children),
	)
}