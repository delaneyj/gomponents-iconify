package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHospitalTwo0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" d="M32 11h8l4 10H4l4-10h8"/><path fill="#fff" stroke="#fff" d="M8 21h32v22H8V21Zm8-16h16v16H16z"/><path fill="#fff" stroke="#000" d="M16 29h8v14h-8zm8 0h8v14h-8z"/><path stroke="#000" d="M21 13h6"/><path stroke="#fff" d="M36 43H12"/><path stroke="#000" d="M24 16v-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHospitalTwo0)"/>`),
		g.Group(children),
	)
}