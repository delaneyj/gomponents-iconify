package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssemblyLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAssemblyLine0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="16" cy="10" r="4" fill="#fff"/><path stroke-linecap="round" stroke-linejoin="round" d="M28 38H13c-4 0-7-2.917-7-7s3-7 7-7h7m0 0h15c4 0 7-2.917 7-7s-3-7-7-7H20M6 10h6m24 28h6"/><circle cx="32" cy="38" r="4" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAssemblyLine0)"/>`),
		g.Group(children),
	)
}