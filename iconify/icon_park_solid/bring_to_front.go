package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BringToFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBringToFront0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#fff" stroke-linejoin="round" d="M41 18H19a1 1 0 0 0-1 1v22a1 1 0 0 0 1 1h22a1 1 0 0 0 1-1V19a1 1 0 0 0-1-1Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M9.97 6H6v4.034M9.997 30H6v-3.988M26.002 6H30v4.015M16.028 6h3.98M6 16v4m24-4v2.007"/><path stroke-linecap="round" d="M15.992 30H18"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBringToFront0)"/>`),
		g.Group(children),
	)
}