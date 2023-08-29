package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagLibya(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5c9e31" d="M5 17h62v38H5z"/><path fill="#d22f27" d="M5 17h62v13H5z"/><path d="M5 30h62v12H5z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M32.136 36.125a4.053 4.053 0 0 1 3.378-3.924a4.395 4.395 0 0 0-.81-.076a4.004 4.004 0 1 0 0 8a4.395 4.395 0 0 0 .81-.076a4.053 4.053 0 0 1-3.378-3.924Zm6.057 1.041l1.328-4l1.145 3.939l-3.18-2.373l4-.099l-3.293 2.533z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}