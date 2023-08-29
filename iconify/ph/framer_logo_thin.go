package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FramerLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204 96V32a4 4 0 0 0-4-4H56a4 4 0 0 0-2.66 7l64.14 57H56a4 4 0 0 0-4 4v64a4 4 0 0 0 1.17 2.83l72 72A4 4 0 0 0 132 232v-68h68a4 4 0 0 0 2.66-7l-64.14-57H200a4 4 0 0 0 4-4Zm-14.52 60H128a4 4 0 0 0-4 4v62.34l-64-64V100h66.48ZM196 92h-66.48l-63-56H196Z"/>`),
		g.Group(children),
	)
}