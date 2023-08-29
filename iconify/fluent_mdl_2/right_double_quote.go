package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightDoubleQuote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 66h332v343q0 51-7 105t-23 104t-41 96t-62 78t-84 53T5 865V708q37 0 64-14t47-38t30-55t18-64t9-66t3-62H0V66zm875 0v343q0 78-16 159t-54 147t-100 108t-157 42V708q55 0 89-30t53-76t25-98t7-95H546V66h329z"/>`),
		g.Group(children),
	)
}