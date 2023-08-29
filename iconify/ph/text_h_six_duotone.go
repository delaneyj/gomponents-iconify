package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHSixDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M240 64v136a8 8 0 0 1-8 8H48a8 8 0 0 1-8-8V56h192a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M152 56v120a8 8 0 0 1-16 0v-52H48v52a8 8 0 0 1-16 0V56a8 8 0 0 1 16 0v52h88V56a8 8 0 0 1 16 0Zm96 124a36 36 0 1 1-67.34-17.68c.07-.14.14-.28.22-.42l32.25-54a8 8 0 0 1 13.74 8.2l-16.69 28c.6 0 1.21-.05 1.82-.05A36 36 0 0 1 248 180Zm-16 0a20 20 0 1 0-20 20a20 20 0 0 0 20-20Z"/></g>`),
		g.Group(children),
	)
}