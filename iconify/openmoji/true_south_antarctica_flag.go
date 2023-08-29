package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrueSouthAntarcticaFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FFF" d="M5 17.08h62v38H5z"/><path fill="#1E50A0" stroke="#1E50A0" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17.08h62v19H46.333L36 51.191L25.667 36.08H5v-19"/><path fill="#FFF" stroke="#FFF" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m36 21.468l9.833 15.111L36 40.357l-9.833-3.778L36 21.468"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}