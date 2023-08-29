package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkiffCalendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.92 3.43C10.508.217 1.712 15.697 3.163 26.996c2.18 16.973 22.342 21.966 25.905 3.945c-16.34-.875-18.032-24.781-.148-27.509V3.43Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.92 3.43c-32.874.453-16.924 47.309.148 27.509"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.08 44.57c18.412 3.213 27.208-12.267 25.757-23.565C42.657 4.031 22.495-.963 18.931 17.06c16.34.875 18.033 24.781.148 27.51h0v.001Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.08 44.57c32.874-.454 16.924-47.31-.148-27.51"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.777 24.824h0a1.77 1.77 0 0 1-1.778 1.778h0c-.986 0-1.776-.791-1.776-1.712v-1.777c0-.986.79-1.777 1.777-1.71h0a1.7 1.7 0 0 1 1.711 1.688v.022h0"/>`),
		g.Group(children),
	)
}