package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceIdIrisScanIdentificationRetinaSecureSecurityIrisScanEyeBrackets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 10.5v2a1 1 0 0 1-1 1h-2m0-13h2a1 1 0 0 1 1 1v2m-13 0v-2a1 1 0 0 1 1-1h2m0 13h-2a1 1 0 0 1-1-1v-2m11-3.5s-2 3-4.5 3s-4.5-3-4.5-3s2-3 4.5-3s4.5 3 4.5 3Z"/><circle cx="7" cy="7" r=".5"/></g>`),
		g.Group(children),
	)
}