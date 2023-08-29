package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewMail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 1280v256h256v128h-256v256h-128v-256h-256v-128h256v-256h128zM0 384h2048v896h-128V648l-896 447l-896-447v888h1024v128H0V384zm143 128l881 441l881-441H143z"/>`),
		g.Group(children),
	)
}