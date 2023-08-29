package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstanceRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m11.414 18.154l4.062-5.803l.1-.148L17.116 10l-1.542-2.203a4.21 4.21 0 0 0-.099-.148l-2.914-4.164l-1.148-1.64a2.002 2.002 0 0 1 2.783.496l4.559 6.512a2 2 0 0 1 0 2.294l-4.559 6.512a2.002 2.002 0 0 1-2.783.495Zm2.342-9.301L9.197 2.341a2 2 0 0 0-3.277 0L1.362 8.853a2 2 0 0 0 0 2.294l4.558 6.512a2 2 0 0 0 3.277 0l4.559-6.512a2 2 0 0 0 0-2.294ZM5.92 17.659l1.639-1.147L3 10l4.56-6.513l1.638-1.146L7.56 3.488L12.117 10L7.56 16.512L5.92 17.66Z"/>`),
		g.Group(children),
	)
}