package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Section(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1620 576q23 0 46 8t44 23t37 34t24 42l277 789H0l278-789q8-22 23-41t36-34t43-23t48-9h1192zM428 704q-3 0-7 2t-10 6t-8 6t-5 7l-217 619h1686l-217-619q-1-3-4-6t-9-7t-9-5t-8-3H428z"/>`),
		g.Group(children),
	)
}