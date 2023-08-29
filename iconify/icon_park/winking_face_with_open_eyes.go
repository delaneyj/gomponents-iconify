package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WinkingFaceWithOpenEyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M24 44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44Z"/><path stroke="#fff" stroke-linecap="round" d="M33 16L29 20L33 24"/><path stroke="#fff" stroke-linecap="round" d="M31 31C31 31 29 35 24 35C19 35 17 31 17 31"/><circle cx="17" cy="20" r="4" fill="#43CCF8" stroke="#fff" stroke-linecap="round"/></g>`),
		g.Group(children),
	)
}