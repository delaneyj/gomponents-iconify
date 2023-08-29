package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BretagneFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fff" d="M5 17h62v38H5z"/><path d="M33 22h34v4H33zm0 8h34v4H33zM5 38h62v4H5zm0 8h62v4H5z"/><path stroke="#000" stroke-linejoin="round" stroke-width="2" d="M9 24h2l-1-3zm9 0h2l-1-3zm9 0h2l-1-3zm-13.5 5h2l-1-3zm9 0h2l-1-3zM9 34h2l-1-3zm9 0h2l-1-3zm9 0h2l-1-3z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}