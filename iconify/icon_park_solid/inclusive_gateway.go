package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InclusiveGateway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSInclusiveGateway0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M22.8 4.201L4.413 22.586a2 2 0 0 0 0 2.828L22.8 43.8a2 2 0 0 0 2.829 0l18.384-18.385a2 2 0 0 0 0-2.828L25.628 4.2a2 2 0 0 0-2.829 0Z"/><path fill="#000" stroke="#000" d="M24 32a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSInclusiveGateway0)"/>`),
		g.Group(children),
	)
}