package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DependencyAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 512v640h512v128H0V384h1280v512h-128V512H128zm640 1408v-896h1280v896H768zm128-768v640h1024v-640H896zm384-1023V0h257q51 0 98 19t84 56q36 36 55 83t20 99v294l164-162l90 90l-318 318l-319-318l91-90l164 162V257q0-27-10-50t-27-41t-41-27t-51-10h-257z"/>`),
		g.Group(children),
	)
}