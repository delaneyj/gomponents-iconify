package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyPants(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M34 4H14C12 5 9 9 9 19C9 31.5 14 44 17 44C19 44 19 41.5 19 41.5V24C19 19 27 19.5 27 24V35C27 43 34 46 36 44C38 42 34 40 34 37C34 32 38 28 38 15C38 10 35.6667 6 34 4Z"/><path stroke="#fff" stroke-linecap="round" d="M10 12H37"/><path stroke="#000" d="M10.8032 8C10.0707 9.69577 9.46541 11.9812 9.17871 15"/><path stroke="#000" d="M37.9995 15C37.9995 12.7719 37.5362 10.7424 36.875 9"/></g>`),
		g.Group(children),
	)
}