package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHTwoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M240 64v144H48a8 8 0 0 1-8-8V56h192a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M248 208a8 8 0 0 1-8 8h-48a8 8 0 0 1-6.4-12.8l43.16-57.56a16 16 0 1 0-25.54-19.27a16.28 16.28 0 0 0-2.32 4.3a8 8 0 1 1-15.08-5.34a32 32 0 1 1 55.73 29.93L208 200h32a8 8 0 0 1 8 8ZM144 48a8 8 0 0 0-8 8v52H48V56a8 8 0 0 0-16 0v120a8 8 0 0 0 16 0v-52h88v52a8 8 0 0 0 16 0V56a8 8 0 0 0-8-8Z"/></g>`),
		g.Group(children),
	)
}