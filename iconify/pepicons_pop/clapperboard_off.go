package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClapperboardOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd"><path d="M.707 6.08a1 1 0 0 1 .763-1.192L16.123 1.68a1 1 0 0 1 1.19.763l.642 2.93a1 1 0 0 1-.763 1.192L2.54 9.773a1 1 0 0 1-1.19-.763L.707 6.08Zm2.168.548l.213.977l12.7-2.78l-.214-.977l-12.7 2.78Z"/><path d="M9.935 7.698L7.339 5.195l1.389-1.44l2.595 2.503l-1.388 1.44Zm-3.908.855L3.432 6.05L4.82 4.61l2.595 2.504l-1.388 1.44Zm7.815-1.711L11.247 4.34l1.388-1.44l2.595 2.504l-1.388 1.44Zm-4.01 5.713l2-3l-1.664-1.11l-2 3l1.664 1.11Zm4 0l2-3l-1.664-1.11l-2 3l1.664 1.11Zm-8 0l2-3l-1.664-1.11l-2 3l1.664 1.11Z"/><path d="M1.5 9a1 1 0 0 1 1-1h15a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-15a1 1 0 0 1-1-1V9Zm2 1v7h13v-7h-13Z"/><path d="M18 13H2v-2h16v2Z"/></g><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}