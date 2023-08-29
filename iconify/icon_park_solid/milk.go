package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Milk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMilk0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M12 19.574a2 2 0 0 1 .304-1.06L17 11h14l4.696 7.514a2 2 0 0 1 .304 1.06V42a2 2 0 0 1-2 2H14a2 2 0 0 1-2-2V19.574Z"/><path stroke="#000" d="M19 33v-9l5 6l5-6v9"/><path stroke="#fff" d="M17 4h14v7H17V4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMilk0)"/>`),
		g.Group(children),
	)
}