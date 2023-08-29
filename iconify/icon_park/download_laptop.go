package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadLaptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M25 9H11C9.34315 9 8 10.3431 8 12V33H40V25"/><path fill="#2F88FF" d="M4 33H44V35C44 38.3137 41.3137 41 38 41H10C6.68629 41 4 38.3137 4 35V33Z"/><path stroke-linecap="round" d="M31 14L36 19L41 14"/><path stroke-linecap="round" d="M36 7L36 19"/></g>`),
		g.Group(children),
	)
}