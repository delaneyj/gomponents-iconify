package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BbcWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 24.165h8.801V42.5H5.5zm18.666-9.864V5.5h18.335v8.801zm-6.807 9.31l6.223-6.223l12.965 12.965l-6.224 6.223z"/>`),
		g.Group(children),
	)
}