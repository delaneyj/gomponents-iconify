package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighVoltage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#fbb11c" d="M32 29L42 1L12 35h20L22 63l30-34"/>`),
		g.Group(children),
	)
}