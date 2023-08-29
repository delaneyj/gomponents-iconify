package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagCentralAfricanRepublic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#f1b31c" d="M5 17h62v38H5z"/><path fill="#1e50a0" d="M5 17h62v10H5z"/><path fill="#fff" d="M5 27h62v9H5z"/><path fill="#5c9e31" d="M5 36h62v9H5z"/><path fill="#d22f27" d="M33 17h6v38h-6z"/><path fill="#f1b31c" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" d="m12.906 19.603l1.545 5l-4.045-3.09h5l-4.045 3.09l1.545-5z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}