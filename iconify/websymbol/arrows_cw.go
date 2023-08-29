package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsCw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M707 150L857 0v405H452l153-153q-77-62-176-62q-111 0-193.5 76T144 452H1q9-170 133-287T429 48q157 0 278 102zM429 810q111 0 193-76t91-186h143q-9 170-133 287T429 952q-158 0-279-102L0 1000V595h405L252 748q78 62 177 62z"/>`),
		g.Group(children),
	)
}