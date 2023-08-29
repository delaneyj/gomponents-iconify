package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHospital0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M33 5H5v38h28V5Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M33 21h10v22H33"/><path stroke-linecap="round" d="M13 21h12m-6-6v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHospital0)"/>`),
		g.Group(children),
	)
}