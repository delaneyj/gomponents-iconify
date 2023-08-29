package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerEnviroment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M768 384h512v128H768V384zm0 768h512v128H768v-128zm0 256h512v128H768v-128zm1170 640H110l160-640h242V256q0-26 10-49t27-41t41-28t50-10h768q26 0 49 10t41 27t28 41t10 50v1152h242l160 640zM640 1664h768V256H640v1408zm1134 256l-96-384h-142v256H512v-256H370l-96 384h1500z"/>`),
		g.Group(children),
	)
}