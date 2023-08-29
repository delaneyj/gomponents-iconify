package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slippers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSlippers0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 29h40v6H4v-6Z"/><path fill="#555" d="M23.53 13c-3.5 4-3.5 12-3.5 16h20v-8.5S28 15 23.53 13Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSlippers0)"/>`),
		g.Group(children),
	)
}