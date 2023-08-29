package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cocktails(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 64q0 26-19 45l-749 750v1061h320q26 0 45 19t19 45q0 26-19 45t-45 19H576q-26 0-45-19t-19-45q0-26 19-45t45-19h320V859L147 109q-19-19-19-45t19-45t45-19h1536q26 0 45 19t19 45zM346 128l614 614l613-614H346z"/>`),
		g.Group(children),
	)
}