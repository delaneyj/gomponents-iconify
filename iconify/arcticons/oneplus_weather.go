package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneplusWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.506 15.032a11.25 11.25 0 1 1 17.629 13.65"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.484 18.734a11.591 11.591 0 1 1 8.197 19.787m-8.295.072a9.895 9.895 0 0 1 0-19.79"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.64 26.147a6.136 6.136 0 0 1 6.137 6.136h0a6.136 6.136 0 0 1-6.136 6.137h0m-.001 0l-20.254.172"/>`),
		g.Group(children),
	)
}