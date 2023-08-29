package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagPanama(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fff" d="M5 17h62v38H5z"/><path fill="#1e50a0" d="M5 36h31v19H5z"/><path fill="#d22f27" d="M36 17h31v19H36z"/><path fill="#1e50a0" stroke="#1e50a0" stroke-linecap="round" stroke-linejoin="round" d="m18.962 29.167l1.659-5l1.431 4.924l-3.974-2.967l5-.123l-4.116 3.166z"/><path fill="#d22f27" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" d="m50.962 48.167l1.659-5l1.431 4.924l-3.974-2.967l5-.123l-4.116 3.166z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}