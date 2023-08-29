package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagTurkey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m40.64 33.05l3.052 4.019l-4.934-1.532l4.932-1.541l-3.046 4.025l-.004-4.972M31.29 44.64a8.643 8.643 0 1 1 3.958-16.34a11 11 0 1 0 0 15.38a8.715 8.715 0 0 1-3.958.95z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}