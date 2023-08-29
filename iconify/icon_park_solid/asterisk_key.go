package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsteriskKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAsteriskKey0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><path stroke="#000" d="M24 16v16m-6.553-12.589l13.107 9.178m-.001-9.178L17.447 28.59"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAsteriskKey0)"/>`),
		g.Group(children),
	)
}