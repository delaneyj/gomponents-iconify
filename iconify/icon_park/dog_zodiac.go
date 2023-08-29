package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DogZodiac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 44V19C6 14 9.6 9.6 20 4V13H27V19"/><path d="M16 25C20.0133 26.7807 27.354 30.1237 29 40C29.5 43 35 47 41 40C42.9943 37.8639 43.321 34.3488 37.7642 32.5681"/><path d="M28 36.0005C24.6667 35.6227 17 37.0003 17 44"/></g>`),
		g.Group(children),
	)
}