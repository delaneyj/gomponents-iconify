package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comments(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1664 64q49 0 88.5 39.5T1792 192v512q0 52-38 90t-90 38h-128v384l-388-384H768q-50 0-89-39t-39-89V192q0-50 39-89t89-39h896zM512 704q1 97 80 176q76 76 176 80h384v128q0 49-37.5 88.5T1028 1216H644l-388 384v-384H128q-50 0-89-39t-39-89V576q0-50 39-89t89-39h384v256z"/>`),
		g.Group(children),
	)
}