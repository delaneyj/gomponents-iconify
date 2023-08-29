package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLocalPin0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M24 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z"/><path stroke-linecap="round" d="M24 20v18m-8-6h-4L4 44h40l-8-12h-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLocalPin0)"/>`),
		g.Group(children),
	)
}