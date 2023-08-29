package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Urecord(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.89 13.41v14a7.11 7.11 0 1 0 14.22 0v-14"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M31.11 27.69a5.43 5.43 0 0 1 5.54-5.28h0m-5.54 0v14"/>`),
		g.Group(children),
	)
}