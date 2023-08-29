package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 192q115 0 221 29t199 84t168 130t130 168t84 199t30 222q0 115-29 221t-84 199t-130 168t-168 130t-199 84t-222 30q-115 0-221-29t-199-84t-168-130t-130-168t-84-199t-30-222q0-115 29-221t84-199t130-168t168-130t199-84t222-30z"/>`),
		g.Group(children),
	)
}