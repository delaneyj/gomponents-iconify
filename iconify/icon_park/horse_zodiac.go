package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorseZodiac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M18 4C13 4 4 8 4 18.2979V25"/><path stroke-linejoin="round" d="M35 30V44"/><path stroke-linejoin="round" d="M18.0137 22C18.5001 23.5 20.0001 26.5 24.3963 25.7384C26.9798 25.4269 33.0823 26.14 36.0001 31C37.5001 33.5 41.4471 33.4957 44.0001 27.1403"/><path stroke-linejoin="round" d="M26 16.0003L24 13.5004C24 13.5004 22.3833 11.7447 21 11.0005C19.6167 10.2562 17 9.50025 14 11.0004C12.7027 11.8147 10 13.0003 10 17.9427V44"/><path d="M27 44V41C27 38.2386 24.7614 36 22 36V36C19.2386 36 17 38.2386 17 41V44"/></g>`),
		g.Group(children),
	)
}