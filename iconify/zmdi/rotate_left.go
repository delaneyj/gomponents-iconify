package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M65 166q-17 24-22 53H0q6-46 35-83zm-22 95q5 28 22 53l-30 30Q6 307 0 261h43zm22 114l30-31q24 17 53 22v43q-46-5-83-34zM191 71q63 8 106 56t43 113t-43 113t-106 56v-43q45-8 75.5-43.5T297 240t-30.5-82.5T191 114v83l-98-95l98-97v66z"/>`),
		g.Group(children),
	)
}