package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHOneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M224 64v144H48a8 8 0 0 1-8-8V56h176a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M152 56v120a8 8 0 0 1-16 0v-52H48v52a8 8 0 0 1-16 0V56a8 8 0 0 1 16 0v52h88V56a8 8 0 0 1 16 0Zm75.77 49a8 8 0 0 0-8.21.39l-24 16a8 8 0 1 0 8.88 13.32L216 127v81a8 8 0 0 0 16 0v-96a8 8 0 0 0-4.23-7Z"/></g>`),
		g.Group(children),
	)
}