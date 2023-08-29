package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eraser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m896 1152l336-384H464l-336 384h768zM1909 75q15 34 9.5 71.5T1888 212L992 1236q-38 44-96 44H128q-38 0-69.5-20.5T11 1205q-15-34-9.5-71.5T32 1068L928 44q38-44 96-44h768q38 0 69.5 20.5T1909 75z"/>`),
		g.Group(children),
	)
}