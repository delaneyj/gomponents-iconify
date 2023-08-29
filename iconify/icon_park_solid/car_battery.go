package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarBattery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCarBattery0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M43 12H3v30h40V12ZM11 6H3v6h8V6Zm32 0h-8v6h8V6Z"/><path stroke="#000" d="M9 21h6m16 0h6m-25-3v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCarBattery0)"/>`),
		g.Group(children),
	)
}