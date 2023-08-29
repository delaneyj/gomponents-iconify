package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HairDryer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#000" d="M19.193 21.5439L21.512 40.0964C21.7711 42.1692 20.1548 44.0001 18.0659 44.0001C16.3144 44.0001 14.837 42.6959 14.6197 40.9579L12.2461 21.969"/><path fill="#2F88FF" stroke="#000" d="M13 4C8.02944 4 4 8.02944 4 13C4 17.7167 7.62831 21.5859 12.2461 21.9689C12.4947 21.9895 12.7461 22 13 22C14.578 22 16.7395 21.8249 19.193 21.5438L31.5965 19.5219L44 17.5V8.5L28.5 6.25L13 4Z"/><path stroke="#fff" d="M37 8.20068V17.7996"/><path stroke="#000" d="M44.0002 17.5L31.5967 19.5219"/><path stroke="#000" d="M44 8.5L28.5 6.25"/><path fill="#43CCF8" stroke="#fff" d="M16 13C16 14.6569 14.6569 16 13 16C11.3431 16 10 14.6569 10 13C10 11.3431 11.3431 10 13 10C14.6569 10 16 11.3431 16 13Z"/></g>`),
		g.Group(children),
	)
}