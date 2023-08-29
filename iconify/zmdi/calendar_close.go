package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarClose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m135 347l-23-23l52-52l-52-52l23-23l52 52l52-52l22 23l-52 52l52 52l-22 23l-52-52zM341 48q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48h21V5h43v43h170V5h43v43h21zm0 341V155H43v234h298z"/>`),
		g.Group(children),
	)
}