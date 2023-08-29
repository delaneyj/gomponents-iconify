package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Other(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSOther0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M22.799 4.201L4.414 22.586a2 2 0 0 0 0 2.828L22.8 43.8a2 2 0 0 0 2.828 0l18.385-18.385a2 2 0 0 0 0-2.828L25.627 4.2a2 2 0 0 0-2.828 0Z"/><path stroke="#000" stroke-linecap="round" d="M18 24h12m-6-6v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSOther0)"/>`),
		g.Group(children),
	)
}