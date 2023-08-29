package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagSyria(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M5 17h62v38H5z"/><path fill="#d22f27" d="M5 17h62v13H5z"/><path fill="#fff" d="M5 30h62v12H5z"/><path fill="#5c9e31" stroke="#5c9e31" stroke-linecap="round" stroke-linejoin="round" d="m28.5 33.59l1.545 5L26 35.5h5l-4.045 3.09l1.545-5zm15 0l1.545 5L41 35.5h5l-4.045 3.09l1.545-5z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}