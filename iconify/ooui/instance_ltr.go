package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstanceLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.703 1.845L4.641 7.65a4.115 4.115 0 0 0-.099.148L3 10l1.542 2.203c.032.05.065.1.1.148l2.914 4.164l1.147 1.64a2.002 2.002 0 0 1-2.783-.496l-4.558-6.512a2 2 0 0 1 0-2.294L5.92 2.341a2.002 2.002 0 0 1 2.783-.496Zm-2.341 9.302l4.558 6.512a2 2 0 0 0 3.277 0l4.559-6.512a2 2 0 0 0 0-2.294l-4.559-6.512a2 2 0 0 0-3.277 0L6.362 8.853a2 2 0 0 0 0 2.294Zm7.835-8.806L12.56 3.488L17.117 10l-4.558 6.512l-1.639 1.147l1.639-1.147L8 10l4.559-6.512l1.638-1.147Z"/>`),
		g.Group(children),
	)
}