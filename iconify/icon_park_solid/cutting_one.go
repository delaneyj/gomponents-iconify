package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CuttingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCuttingOne0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#fff" stroke-linejoin="round" d="M11 42a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm26 0a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" d="m15.377 39.413l2.123-3.597l17-29.445"/><path stroke-linecap="round" d="m13.496 6.175l17 29.445l2.13 3.793"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCuttingOne0)"/>`),
		g.Group(children),
	)
}