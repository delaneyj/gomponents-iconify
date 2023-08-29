package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubstitutionsIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1045 0l875 1278h-512v770H638v-770H128L1045 0zm235 1150h397l-635-927l-665 927h389v770h514v-770z"/>`),
		g.Group(children),
	)
}