package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceIdUserIdentificationAngleSecureHumanIdPersonFaceSilhouetteSecurityBrackets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 10.5v2a1 1 0 0 1-1 1h-2m0-13h2a1 1 0 0 1 1 1v2m-13 0v-2a1 1 0 0 1 1-1h2m0 13h-2a1 1 0 0 1-1-1v-2"/><circle cx="7" cy="4.5" r="2"/><path d="M10.16 10.5a3.5 3.5 0 0 0-6.32 0"/></g>`),
		g.Group(children),
	)
}