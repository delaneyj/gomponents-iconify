package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Attach(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 384v1216q0 93-35 174t-96 142t-142 96t-175 36q-93 0-174-35t-142-96t-96-142t-36-175V320q0-66 25-124t69-101t102-69T960 0q66 0 124 25t101 69t69 102t26 124v1280q0 40-15 75t-41 61t-61 41t-75 15q-40 0-75-15t-61-41t-41-61t-15-75V512h128v1088q0 26 19 45t45 19q26 0 45-19t19-45V320q0-40-15-75t-41-61t-61-41t-75-15q-40 0-75 15t-61 41t-41 61t-15 75v1280q0 66 25 124t69 101t102 69t124 26q66 0 124-25t101-69t69-102t26-124V384h128z"/>`),
		g.Group(children),
	)
}