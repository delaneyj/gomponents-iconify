package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m14.5.5l.457.203A.5.5 0 0 0 14.5 0v.5ZM.5.5V0a.5.5 0 0 0-.5.5h.5Zm14 9v.5a.5.5 0 0 0 .457-.703L14.5 9.5Zm-2-4.5l-.457-.203a.5.5 0 0 0 0 .406L12.5 5Zm2-5H.5v1h14V0ZM0 .5v9h1v-9H0ZM.5 10h14V9H.5v1Zm14.457-.703l-2-4.5l-.914.406l2 4.5l.914-.406Zm-2-4.094l2-4.5l-.914-.406l-2 4.5l.914.406ZM1 15V9.5H0V15h1Z"/>`),
		g.Group(children),
	)
}