package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 368h-86v43h43l-64 64l-64-64h43v-43H128q-18 0-30.5-12.5T85 325V155H0v-43h85V69H43l64-64l64 64h-43v256h299v43zM171 155v-43h128q17 0 29.5 12.5T341 155v128h-42V155H171z"/>`),
		g.Group(children),
	)
}