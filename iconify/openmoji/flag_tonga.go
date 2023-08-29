package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagTonga(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><path fill="#fff" d="M5 17h32v20H5z"/><path fill="#d22f27" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M27.5 26h-5v-5h-3v5h-5v3h5v5h3v-5h5v-3z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}