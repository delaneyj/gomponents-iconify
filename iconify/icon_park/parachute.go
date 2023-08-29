package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parachute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M16.7273 24C19.1515 21.5758 21.5758 20.3636 24 20.3636C26.4242 20.3636 28.8485 21.5758 31.2727 24C34.101 21.5758 36.2222 20.3636 37.6364 20.3636C39.0505 20.3636 41.1717 21.5758 44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24C6.82828 21.5758 8.9495 20.3636 10.3636 20.3636C11.7778 20.3636 13.899 21.5758 16.7273 24Z"/><path stroke-linecap="round" d="M4 24L24 44L16.7273 24"/><path stroke-linecap="round" d="M31.2727 24L24 44L44 24"/></g>`),
		g.Group(children),
	)
}