package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Indent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M352 704q0 14-9 23L55 1015q-9 9-23 9q-13 0-22.5-9.5T0 992V416q0-13 9.5-22.5T32 384q14 0 23 9l288 288q9 9 9 23zm1440 480v192q0 13-9.5 22.5t-22.5 9.5H32q-13 0-22.5-9.5T0 1376v-192q0-13 9.5-22.5T32 1152h1728q13 0 22.5 9.5t9.5 22.5zm0-384v192q0 13-9.5 22.5t-22.5 9.5H672q-13 0-22.5-9.5T640 992V800q0-13 9.5-22.5T672 768h1088q13 0 22.5 9.5t9.5 22.5zm0-384v192q0 13-9.5 22.5T1760 640H672q-13 0-22.5-9.5T640 608V416q0-13 9.5-22.5T672 384h1088q13 0 22.5 9.5t9.5 22.5zm0-384v192q0 13-9.5 22.5T1760 256H32q-13 0-22.5-9.5T0 224V32Q0 19 9.5 9.5T32 0h1728q13 0 22.5 9.5T1792 32z"/>`),
		g.Group(children),
	)
}