package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outbound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M40 24L24 24"/><path stroke-linecap="round" stroke-linejoin="round" d="M34 18L40 24L34 30"/><circle cx="20" cy="24" r="4" fill="#2F88FF"/><path stroke-linecap="round" d="M40.706 13C39.9214 11.8109 39.0133 10.7105 38 9.71713C34.3925 6.18064 29.4509 4 24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44C29.4509 44 34.3925 41.8194 38 38.2829C39.0133 37.2895 39.9214 36.1891 40.706 35"/></g>`),
		g.Group(children),
	)
}