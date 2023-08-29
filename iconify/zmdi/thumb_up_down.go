package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbUpDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 128v27q0 6-2 11l-49 113q-8 20-29 20H32q-13 0-22.5-9.5T0 267V128q0-13 9-23L115 0l17 17q7 7 7 17l-1 5l-14 68h111q8 0 14.5 6t6.5 15zm224 85q13 0 22.5 9.5T512 245v139q0 13-9 23L397 512l-17-17q-7-7-7-17l1-5l14-68H277q-8 0-14.5-6t-6.5-15v-27q0-6 2-11l49-113q8-20 29-20h144z"/>`),
		g.Group(children),
	)
}