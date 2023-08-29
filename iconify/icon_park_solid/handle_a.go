package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandleA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHandleA0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20" fill="#fff" stroke="#fff"/><path stroke="#000" d="M24 11L14 33m4-7h12m-6-15l10 22"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHandleA0)"/>`),
		g.Group(children),
	)
}