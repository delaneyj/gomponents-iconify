package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cifsdocumentsprovider(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.097 8.44a1.215 1.215 0 0 0-1.258-1.172H26.516C25.161 7.202 23.644 4.5 22.113 4.5h-7.959a1.459 1.459 0 0 0-1.476 1.475v17.2a1.459 1.459 0 0 0 1.47 1.454H37.62a1.469 1.469 0 0 0 1.477-1.454ZM8.903 35.323h10.065m-5.033-5.662v5.662m14.326-.339l6.662-3.79m-6.633 5.672l6.618 3.88"/><circle cx="26.642" cy="35.952" r="1.887" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.581" cy="30.29" r="1.887" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.581" cy="41.613" r="1.887" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}