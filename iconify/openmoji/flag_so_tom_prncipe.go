package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagSoTomPrncipe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#b1cc33" d="M5 17h62v38H5z"/><path fill="#fcea2b" d="M5 30h62v12H5z"/><path fill="#d22f27" d="M26 36L5 55V17l21 19z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M38.5 33.59l1.545 5L36 35.5h5l-4.045 3.09l1.545-5z" fill="#000"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M53.5 33.59l1.545 5L51 35.5h5l-4.045 3.09l1.545-5z" fill="#000"/><g><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/></g>`),
		g.Group(children),
	)
}