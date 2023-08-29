package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M291.75 308.1a105.8 105.8 0 1 0 0-211.6H136v32h39.943v147.6H136v32h39.943V352H136v32h39.943v56h32v-56H304v-32h-96.057v-43.9Zm-83.807-179.6h83.807a73.8 73.8 0 1 1 0 147.6h-83.807Z"/>`),
		g.Group(children),
	)
}