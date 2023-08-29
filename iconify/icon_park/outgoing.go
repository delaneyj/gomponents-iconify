package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outgoing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path fill="#2F88FF" stroke-linejoin="round" d="M18 35C18 32.7909 16.2091 31 14 31C11.7909 31 10 32.7909 10 35C10 37.2091 11.7909 39 14 39C16.2091 39 18 37.2091 18 35Z"/><path fill="#2F88FF" stroke-linejoin="round" d="M37 35C37 32.7909 35.2091 31 33 31C30.7909 31 29 32.7909 29 35C29 37.2091 30.7909 39 33 39C35.2091 39 37 37.2091 37 35Z"/><path stroke-linecap="round" d="M4 35H10"/><path stroke-linecap="round" d="M18 35H29"/><path stroke-linecap="round" d="M37 35H44"/><path stroke-linecap="round" stroke-linejoin="round" d="M38 19L44 13L38 7"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 13H44"/></g>`),
		g.Group(children),
	)
}