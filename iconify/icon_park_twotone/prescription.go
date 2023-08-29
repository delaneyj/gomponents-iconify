package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Prescription(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPrescription0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" fill-rule="evenodd" stroke-linejoin="round" d="M8 8a2 2 0 0 1 2-2h17v12h13v22a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2V8Z" clip-rule="evenodd"/><path stroke-linecap="round" stroke-linejoin="round" d="m27 6l13 12M27.024 6v12.082H40"/><path stroke-linecap="round" d="M14 30h12m-6-6v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPrescription0)"/>`),
		g.Group(children),
	)
}