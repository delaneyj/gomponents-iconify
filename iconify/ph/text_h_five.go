package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M152 56v120a8 8 0 0 1-16 0v-52H48v52a8 8 0 0 1-16 0V56a8 8 0 0 1 16 0v52h88V56a8 8 0 0 1 16 0Zm60 88a38.8 38.8 0 0 0-9.41 1.14l4.19-25.14H240a8 8 0 0 0 0-16h-40a8 8 0 0 0-7.89 6.68l-8 48a8 8 0 0 0 13.6 6.92A19.73 19.73 0 0 1 212 160a20 20 0 0 1 0 40a19.73 19.73 0 0 1-14.29-5.6a8 8 0 1 0-11.42 11.2A35.54 35.54 0 0 0 212 216a36 36 0 0 0 0-72Z"/>`),
		g.Group(children),
	)
}