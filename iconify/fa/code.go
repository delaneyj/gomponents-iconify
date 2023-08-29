package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m585 1143l-50 50q-10 10-23 10t-23-10L23 727q-10-10-10-23t10-23l466-466q10-10 23-10t23 10l50 50q10 10 10 23t-10 23L192 704l393 393q10 10 10 23t-10 23zM1176 76L803 1367q-4 13-15.5 19.5T764 1389l-62-17q-13-4-19.5-15.5T680 1332L1053 41q4-13 15.5-19.5T1092 19l62 17q13 4 19.5 15.5T1176 76zm657 651l-466 466q-10 10-23 10t-23-10l-50-50q-10-10-10-23t10-23l393-393l-393-393q-10-10-10-23t10-23l50-50q10-10 23-10t23 10l466 466q10 10 10 23t-10 23z"/>`),
		g.Group(children),
	)
}