package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoiceOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M24 44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44Z"/><path stroke="#fff" stroke-linecap="round" d="M30 18V30"/><path stroke="#fff" stroke-linecap="round" d="M36 22V26"/><path stroke="#fff" stroke-linecap="round" d="M18 18V30"/><path stroke="#fff" stroke-linecap="round" d="M12 22V26"/><path stroke="#fff" stroke-linecap="round" d="M24 14V34"/></g>`),
		g.Group(children),
	)
}