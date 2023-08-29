package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagIsrael(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fff" d="M5 17h62v38H5z"/><path fill="#1e50a0" d="M5 20h62v3H5zm0 29h62v3H5z"/><path fill="none" stroke="#1e50a0" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m36 26l-9 15h18l-9-15z"/><path fill="none" stroke="#1e50a0" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m36 46l-9-15h18l-9 15z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}