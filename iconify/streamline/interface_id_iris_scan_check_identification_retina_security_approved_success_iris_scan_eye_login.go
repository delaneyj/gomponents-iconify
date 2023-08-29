package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceIdIrisScanCheckIdentificationRetinaSecurityApprovedSuccessIrisScanEyeLogin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m4.5 7.25l1.55 1.14a.48.48 0 0 0 .4.1a.5.5 0 0 0 .34-.24l3.71-6"/><path d="M11.85 5.25a13.92 13.92 0 0 1 1.65 2s-2.91 4.5-6.5 4.5S.5 7.25.5 7.25s2.91-4.5 6.5-4.5"/></g>`),
		g.Group(children),
	)
}