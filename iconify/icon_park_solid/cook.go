package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCook0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M42 36V20H14v16a6 6 0 0 0 6 6h16a6 6 0 0 0 6-6Z"/><path d="M4 20h40M18 8v4m10-6v6m10-4v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCook0)"/>`),
		g.Group(children),
	)
}