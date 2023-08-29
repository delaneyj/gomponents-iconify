package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullScreenPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFullScreenPlay0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M16 40H6a2 2 0 0 1-2-2V10a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v6"/><path fill="#555" d="M42 24H26a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V26a2 2 0 0 0-2-2Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFullScreenPlay0)"/>`),
		g.Group(children),
	)
}