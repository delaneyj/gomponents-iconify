package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17.5 8.5a2 2 0 0 0-2-2.003l-5 .003l-2-2h-4a1 1 0 0 0-1 1v3"/><path d="m2.81 9.742l1.311 5.243a2 2 0 0 0 1.94 1.515h8.877a2 2 0 0 0 1.94-1.515L18.19 9.74a1 1 0 0 0-.97-1.243L3.781 8.5a1 1 0 0 0-.97 1.243z"/></g>`),
		g.Group(children),
	)
}