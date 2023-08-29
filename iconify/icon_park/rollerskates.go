package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rollerskates(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path stroke="#000" d="M44 4H28V12H44V4Z"/><path fill="#2F88FF" stroke="#000" d="M44 12V34C44 35.11 43.11 36 42 36H8C5.79 36 4 34.21 4 32V26C4 21.58 7.58 18 12 18H28V12H44Z"/><path stroke="#fff" d="M14 24V18"/><path stroke="#fff" d="M21 24V18"/><path stroke="#000" d="M23 18L12 18"/><path stroke="#000" d="M9 44C11.2091 44 13 42.2091 13 40C13 37.7909 11.2091 36 9 36C6.79086 36 5 37.7909 5 40C5 42.2091 6.79086 44 9 44Z"/><path stroke="#000" d="M19.3301 44C21.5392 44 23.3301 42.2091 23.3301 40C23.3301 37.7909 21.5392 36 19.3301 36C17.1209 36 15.3301 37.7909 15.3301 40C15.3301 42.2091 17.1209 44 19.3301 44Z"/><path stroke="#000" d="M29.6699 44C31.8791 44 33.6699 42.2091 33.6699 40C33.6699 37.7909 31.8791 36 29.6699 36C27.4608 36 25.6699 37.7909 25.6699 40C25.6699 42.2091 27.4608 44 29.6699 44Z"/><path stroke="#000" d="M40 44C42.2091 44 44 42.2091 44 40C44 37.7909 42.2091 36 40 36C37.7909 36 36 37.7909 36 40C36 42.2091 37.7909 44 40 44Z"/></g>`),
		g.Group(children),
	)
}