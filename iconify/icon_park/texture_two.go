package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextureTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M40 4H8C5.79086 4 4 5.79086 4 8V40C4 42.2091 5.79086 44 8 44H40C42.2091 44 44 42.2091 44 40V8C44 5.79086 42.2091 4 40 4Z"/><path stroke="#fff" d="M12 12V20"/><path stroke="#fff" d="M28 28V36"/><path stroke="#fff" d="M20 12V20"/><path stroke="#fff" d="M28 12H36"/><path stroke="#fff" d="M12 28H20"/><path stroke="#fff" d="M28 20H36"/><path stroke="#fff" d="M12 36H20"/><path stroke="#fff" d="M36 28V36"/></g>`),
		g.Group(children),
	)
}