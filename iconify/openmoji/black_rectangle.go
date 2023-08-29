package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackRectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M67 17H5v38h62V17Z"/><path fill="#3F3F3F" d="M67 17H5v38h62V17Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M67 17H5v38h62V17Z"/>`),
		g.Group(children),
	)
}