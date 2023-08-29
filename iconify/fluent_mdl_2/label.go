package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Label(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1720 128q41 0 77 16t63 43t43 63t16 78q0 37-10 65t-29 53t-41 46t-47 46v1254H256V256h1254q24-24 46-47t47-41t52-29t65-11zm0 128q-29 0-50 21l-948 948l-34 135l135-34l948-948q21-21 21-50q0-30-21-51t-51-21zm-56 410l-775 776l-377 94l94-377l776-775H384v1280h1280V666z"/>`),
		g.Group(children),
	)
}