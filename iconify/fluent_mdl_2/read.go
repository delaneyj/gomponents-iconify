package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Read(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1024 120l1024 512v1288H0V632l1024-512zm873 580l-873-436l-873 436l324 324h1098l324-324zM128 1792h1792V859l-293 293H421L128 859v933z"/>`),
		g.Group(children),
	)
}