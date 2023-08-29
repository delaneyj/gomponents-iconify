package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagJamaica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5c9e31" d="M5 17h62v38H5z"/><path d="M36 36L5 55V17l31 19zm0 0l31 19V17L36 36z"/><path fill="#fcea2b" d="M5 23v-6h7l55 32v6h-7L5 23z"/><path fill="#fcea2b" d="M67 23v-6h-7L5 49v6h7l55-32z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}