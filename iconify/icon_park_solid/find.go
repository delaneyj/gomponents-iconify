package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Find(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFind0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" d="M4 7h40M4 23h11M4 39h11"/><path fill="#fff" d="M31.5 34a8.5 8.5 0 1 0 0-17a8.5 8.5 0 0 0 0 17Z"/><path stroke-linecap="round" d="m37 32l7 7.05"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFind0)"/>`),
		g.Group(children),
	)
}