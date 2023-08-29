package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagCuracao(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#1e50a0" d="M5 17h62v38H5z"/><path fill="#f1b31c" d="M5 41h62v6H5z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m17.5 26l1.545 5L15 27.91h5L15.955 31l1.545-5zm-6-5l1.545 5L9 22.91h5L9.955 26l1.545-5z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}