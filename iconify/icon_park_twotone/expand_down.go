package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTExpandDown0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M6 10a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v28a4 4 0 0 1-4 4H10a4 4 0 0 1-4-4V10Z"/><path stroke-linecap="round" d="M6 32h36M20 16l4 4l4-4M6 26v12m36-12v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTExpandDown0)"/>`),
		g.Group(children),
	)
}