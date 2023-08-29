package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagLiberia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fff" d="M5 17h62v38H5z"/><path fill="#d22f27" d="M5 17h62v5H5zm0 9h62v4H5zm0 8h62v4H5zm0 8h62v4H5zm0 8h62v5H5z"/><path fill="#1e50a0" d="M5 17h17v17H5z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M14.455 26.09L16 24.91h-1.91L13.5 23l-.59 1.91H11l1.545 1.18l-.59 1.91l1.545-1.18L15.045 28l-.59-1.91z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}