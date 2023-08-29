package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShortSkirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M12 9L17 4H31L36 9L43 35C43 35 39.9997 44 24 44C8.00031 44 5 35 5 35L12 9Z"/><path stroke="#fff" d="M13 42L17 26"/><path stroke="#000" d="M5 35C5 35 8.00031 44 24 44"/></g>`),
		g.Group(children),
	)
}