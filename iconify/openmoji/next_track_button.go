package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextTrackButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="m29.484 17l16.976 9.111l14.286 7.668c1.672.897 1.672 3.545 0 4.442L46.46 45.89L29.484 55"/><path d="m6 17l17.055 9.111l14.353 7.668c1.68.897 1.68 3.545 0 4.442L23.055 45.89L6 55m60-38v38"/></g>`),
		g.Group(children),
	)
}