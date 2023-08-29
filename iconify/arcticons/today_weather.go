package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TodayWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.364 45.5l6.359-1.907l-11.446-3.816l15.897-3.179l-20.348-5.087c22.255-1.907 30.521-4.45 30.521-7.63c0-28.613-40.694-27.977-40.694-1.272"/>`),
		g.Group(children),
	)
}