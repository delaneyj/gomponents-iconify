package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.5 10.5A5.5 5.5 0 0 1 12 16m0 0a5.5 5.5 0 0 1-5.5-5.5M12 16v4m-4 0h4m0 0h4m-4-7a2.5 2.5 0 0 1-2.5-2.5v-4a2.5 2.5 0 0 1 5 0v4A2.5 2.5 0 0 1 12 13Z"/>`),
		g.Group(children),
	)
}