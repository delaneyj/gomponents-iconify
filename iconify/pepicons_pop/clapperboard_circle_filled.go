package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClapperboardCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopClapperboardCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M3.707 9.08a1 1 0 0 1 .763-1.192L19.123 4.68a1 1 0 0 1 1.19.763l.642 2.93a1 1 0 0 1-.763 1.192L5.54 12.773a1 1 0 0 1-1.19-.763l-.642-2.93Zm2.168.548l.213.977l12.7-2.78l-.214-.977l-12.7 2.78Z"/><path d="m12.935 10.698l-2.596-2.503l1.389-1.44l2.595 2.503l-1.388 1.44Zm-3.908.855L6.432 9.05L7.82 7.61l2.595 2.504l-1.388 1.44Zm7.815-1.711L14.247 7.34l1.388-1.44l2.595 2.504l-1.388 1.44Zm-4.01 5.713l2-3l-1.664-1.11l-2 3l1.664 1.11Zm4 0l2-3l-1.664-1.11l-2 3l1.664 1.11Zm-8 0l2-3l-1.664-1.11l-2 3l1.664 1.11Z"/><path d="M4.5 12a1 1 0 0 1 1-1h15a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-15a1 1 0 0 1-1-1v-9Zm2 1v7h13v-7h-13Z"/><path d="M21 16H5v-2h16v2Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopClapperboardCircleFilled0)"/></g>`),
		g.Group(children),
	)
}