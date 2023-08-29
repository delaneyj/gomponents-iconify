package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagMozambique(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fcea2b" d="M5 17h62v38H5z"/><path fill="#186648" d="M5 17h62v13H5z"/><path stroke="#fff" stroke-miterlimit="10" stroke-width="2" d="M5 30h62v12H5z"/><path fill="#d22f27" d="M26 36L5 55V17l21 19z"/><path fill="#fcea2b" stroke="#fcea2b" stroke-linecap="round" stroke-linejoin="round" d="m10.004 41.409l3.589-10.818l3.096 10.654l-8.598-6.42l10.818-.266l-8.905 6.85z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M16.869 38.804h-6.738l.902-3.219h4.934l.902 3.219z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.934 40.121l7.328-7.341m-5.524 0l7.328 7.341"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}