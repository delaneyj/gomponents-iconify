package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Signal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M256 1312v192q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-192q0-14 9-23t23-9h192q14 0 23 9t9 23zm384-128v320q0 14-9 23t-23 9H416q-14 0-23-9t-9-23v-320q0-14 9-23t23-9h192q14 0 23 9t9 23zm384-256v576q0 14-9 23t-23 9H800q-14 0-23-9t-9-23V928q0-14 9-23t23-9h192q14 0 23 9t9 23zm384-384v960q0 14-9 23t-23 9h-192q-14 0-23-9t-9-23V544q0-14 9-23t23-9h192q14 0 23 9t9 23zm384-512v1472q0 14-9 23t-23 9h-192q-14 0-23-9t-9-23V32q0-14 9-23t23-9h192q14 0 23 9t9 23z"/>`),
		g.Group(children),
	)
}