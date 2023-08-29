package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LargeBlueDiamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#92d3f5" d="m12.065 35.5l24.218-24.218L60.5 35.5L36.283 59.718z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12.065 35.5l24.218-24.218L60.5 35.5L36.283 59.718z"/>`),
		g.Group(children),
	)
}