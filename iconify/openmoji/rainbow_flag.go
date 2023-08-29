package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RainbowFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g stroke-width="0"><path fill="#b399c8" d="M5 48h62v6.2H5z"/><path fill="#92d3f5" d="M5 42h62v6H5z"/><path fill="#b1cc33" d="M5 36h62v6H5z"/><path fill="#fcea2b" d="M5 30h62v6H5z"/><path fill="#f4aa41" d="M5 24h62v6H5z"/><path fill="#ea5a47" d="M5 17.8h62V24H5z"/></g><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}