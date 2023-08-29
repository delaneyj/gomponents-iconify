package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCamera0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="m15 12l3-6h12l3 6H15Z"/><rect width="40" height="30" x="4" y="12" fill="#fff" stroke="#fff" rx="3"/><path fill="#000" stroke="#000" d="M24 35a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCamera0)"/>`),
		g.Group(children),
	)
}