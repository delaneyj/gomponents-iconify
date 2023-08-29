package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InclusiveGateway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTInclusiveGateway0"><g fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M22.8 4.201L4.413 22.586a2 2 0 0 0 0 2.828L22.8 43.8a2 2 0 0 0 2.829 0l18.384-18.385a2 2 0 0 0 0-2.828L25.628 4.2a2 2 0 0 0-2.829 0Z"/><path d="M24 32a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTInclusiveGateway0)"/>`),
		g.Group(children),
	)
}