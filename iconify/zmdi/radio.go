package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M26 115L296 5l15 36l-177 71h250q18 0 30.5 12.5T427 155v256q0 17-12.5 29.5T384 453H43q-18 0-30.5-12.5T0 411V155q0-14 7.5-24.5T26 115zm80.5 296q26.5 0 45.5-19t19-45.5t-19-45t-45.5-18.5t-45 18.5t-18.5 45T61.5 392t45 19zM384 240v-85H43v85h256v-43h42v43h43z"/>`),
		g.Group(children),
	)
}