package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPieThree0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20"/><path fill="#555" d="M24 4a20 20 0 0 1 14.58 33.69L24 24V4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPieThree0)"/>`),
		g.Group(children),
	)
}