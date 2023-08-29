package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnergySocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44Z"/><path d="M28 21V16"/><path d="M20 21V16"/><path fill="#2F88FF" d="M24 32C28.4183 32 32 28.4183 32 24V21H16V24C16 28.4183 19.5817 32 24 32Z"/><path d="M24 44V32"/></g>`),
		g.Group(children),
	)
}