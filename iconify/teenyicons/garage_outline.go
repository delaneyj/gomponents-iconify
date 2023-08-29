package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarageOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m.5 5.5l-.29-.407A.5.5 0 0 0 0 5.5h.5Zm7-5l.29-.407a.5.5 0 0 0-.58 0L7.5.5Zm7 5h.5a.5.5 0 0 0-.21-.407l-.29.407Zm-12 2V7H2v.5h.5Zm10 0h.5V7h-.5v.5ZM1 15V5.5H0V15h1ZM.79 5.907l7-5l-.58-.814l-7 5l.58.814Zm6.42-5l7 5l.58-.814l-7-5l-.58.814ZM14 5.5V15h1V5.5h-1ZM3 15V7.5H2V15h1Zm-.5-7h10V7h-10v1Zm9.5-.5V15h1V7.5h-1ZM2.5 11h10v-1h-10v1ZM6 13h3v-1H6v1Z"/>`),
		g.Group(children),
	)
}