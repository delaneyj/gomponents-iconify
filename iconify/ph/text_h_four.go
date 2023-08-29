package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M152 56v120a8 8 0 0 1-16 0v-52H48v52a8 8 0 0 1-16 0V56a8 8 0 0 1 16 0v52h88V56a8 8 0 0 1 16 0Zm96 112h-8v-24a8 8 0 0 0-16 0v24h-28.9l20.49-61.47a8 8 0 0 0-15.18-5.06l-24 72A8 8 0 0 0 184 184h40v24a8 8 0 0 0 16 0v-24h8a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}