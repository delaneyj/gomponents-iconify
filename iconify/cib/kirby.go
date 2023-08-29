package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kirby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 0L1.906 8v16L16 32l14.099-8V8zm6.281 16.339l-3.953 2.13v.292h3.953v3.104H9.718v-3.083h3.943v-.302l-3.953-2.141v-3.818L16 15.912l6.286-3.396v3.823z"/>`),
		g.Group(children),
	)
}