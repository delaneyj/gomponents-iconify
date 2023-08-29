package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Necktie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSNecktie0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m15 36l6-23h6l6 23l-9 8l-9-8Zm6-32h6l3 2l-3 7h-6l-3-7l3-2Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSNecktie0)"/>`),
		g.Group(children),
	)
}