package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarDay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 128h384v1792H0V128h384V0h128v128h1024V0h128v128zm256 1664V640H128v1152h1792zm0-1280V256h-256v128h-128V256H512v128H384V256H128v256h1792zm-256 512v640H384v-640h1280zm-128 512v-384H512v384h1024zm128-768v128H384V768h1280z"/>`),
		g.Group(children),
	)
}