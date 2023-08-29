package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Descend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDescend0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="m11.549 16.625l1.464-5.464s-4.794 1.152-7.1 2.857c-2.308 1.706-2.653 5.399.074 6.973c2.726 1.574 38.186 18.945 38.186 18.945l-2.768-8.794L11.55 16.625Z"/><path d="m20 35l6 3m3-13L26 9l-4-2l-3 13"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDescend0)"/>`),
		g.Group(children),
	)
}