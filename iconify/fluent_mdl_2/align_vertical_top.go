package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 0v128H0V0h2048zm-896 256h512v1152h-512V256zm128 1024h256V384h-256v896zM384 256h512v1664H384V256zm128 1536h256V384H512v1408z"/>`),
		g.Group(children),
	)
}