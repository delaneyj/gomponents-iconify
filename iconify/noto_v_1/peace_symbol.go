package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeaceSymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<g fill="none" stroke="#f79229" stroke-miterlimit="10" stroke-width="11"><circle cx="64" cy="64" r="58"/><path stroke-linecap="round" stroke-linejoin="round" d="M64 6v116m0-58l41.01 41.01m-82.02 0L64 64"/></g>`),
		g.Group(children),
	)
}