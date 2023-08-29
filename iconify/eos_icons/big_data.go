package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BigData(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.965 8a5.35 5.35 0 0 0-3.615-1.96a7.492 7.492 0 0 0-14-2A5.904 5.904 0 0 0 4 4.365A5.98 5.98 0 0 0 4 15.65V8h16v7.9a5.003 5.003 0 0 0 4-4.9a4.908 4.908 0 0 0-1.035-3ZM18.02 22.003c0 1.103-2.687 1.997-6 1.997s-6-.894-6-1.997v-1.996c0 1.102 2.686 1.996 6 1.996s6-.894 6-1.996"/><path fill="currentColor" d="M12.02 16.684c-3.311 0-6-.898-6-1.996v3.964c0 1.099 2.689 1.997 6 1.997s6-.898 6-1.997v-3.964c0 1.098-2.69 1.996-6 1.996Z"/><path fill="currentColor" d="M18.02 11.997c0-1.103-2.687-1.997-6-1.997s-6 .894-6 1.997v1.33c0 1.104 2.686 1.998 6 1.998s6-.894 6-1.997"/>`),
		g.Group(children),
	)
}