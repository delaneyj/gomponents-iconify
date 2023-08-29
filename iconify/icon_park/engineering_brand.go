package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EngineeringBrand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="38" height="20" x="5" y="6" fill="#2F88FF" rx="2"/><path stroke-linecap="round" d="M12 26V42"/><path stroke-linecap="round" d="M36 33H12"/><path stroke-linecap="round" d="M16 42L8 42"/><path stroke-linecap="round" d="M40 42L32 42"/><path stroke-linecap="round" d="M36 26V42"/></g>`),
		g.Group(children),
	)
}