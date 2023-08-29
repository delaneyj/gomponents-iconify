package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Topic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M24 44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24C4 26.7117 4.53967 29.2974 5.51747 31.6554C6.02232 32.8729 6.64396 34.0297 7.36843 35.1119C7.61157 35.4751 7.15543 37.7711 6 42C10.2289 40.8446 12.5249 40.3884 12.8881 40.6316C13.9703 41.356 15.1271 41.9777 16.3446 42.4825C18.7026 43.4603 21.2883 44 24 44Z"/><path stroke="#fff" stroke-linecap="round" d="M16.6045 19.8201H33.3838"/><path stroke="#fff" stroke-linecap="round" d="M21.8467 15.7378L18.933 32.2622"/><path stroke="#fff" stroke-linecap="round" d="M28.8467 15.7378L25.933 32.2622"/><path stroke="#fff" stroke-linecap="round" d="M14.6045 28H31.3838"/></g>`),
		g.Group(children),
	)
}