package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransgenderFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g stroke-width="0"><path fill="#92d3f5" d="M6 46.8h60V54H6z"/><path fill="#ffa7c0" d="M6 39.6h60v7.2H6z"/><path fill="#fff" d="M6 32.4h60v7.2H6z"/><path fill="#ffa7c0" d="M6 25.2h60v7.2H6z"/><path fill="#92d3f5" d="M6 18h60v7.2H6z"/></g><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}