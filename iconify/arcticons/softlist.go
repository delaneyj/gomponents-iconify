package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Softlist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 6.345h6.125l7.398 25.253h19.22L43.5 13.496H18.545"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.12 20.971l4.118 4.117l7.743-7.743"/><circle cx="19.886" cy="38.228" r="3.427" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.241" cy="38.228" r="3.427" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}