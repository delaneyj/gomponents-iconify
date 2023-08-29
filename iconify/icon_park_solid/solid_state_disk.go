package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SolidStateDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSolidStateDisk0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M44 29H4v13h40V29Z"/><path fill="#000" d="M35.5 38a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M4 29L9.038 4.999H39.02l4.98 24"/><path stroke="#000" stroke-linecap="round" stroke-width="4" d="M10 35.5h17"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSolidStateDisk0)"/>`),
		g.Group(children),
	)
}