package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchNintendo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" stroke="#000" stroke-width="4" d="M6 12C6 7.58172 9.58172 4 14 4H19C20.1046 4 21 4.89543 21 6V40C21 41.1046 20.1046 42 19 42H14C9.58172 42 6 38.4183 6 34V12Z"/><path fill="#2F88FF" stroke="#000" stroke-width="4" d="M42 14C42 9.58172 38.4183 6 34 6H28C26.8954 6 26 6.89543 26 8V42C26 43.1046 26.8954 44 28 44H34C38.4183 44 42 40.4183 42 36V14Z"/><rect width="5" height="5" x="11" y="30" fill="#fff" rx="2.5"/><rect width="5" height="5" x="31.5" y="14" fill="#fff" rx="2.5"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M15 17H12"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M34 28V34"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M31 31L37 31"/></g>`),
		g.Group(children),
	)
}