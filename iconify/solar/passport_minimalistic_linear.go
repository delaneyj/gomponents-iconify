package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PassportMinimalisticLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M4 6v13a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H4Zm0 0V5"/><circle cx="12" cy="14" r="3" stroke="currentColor" stroke-width="1.5"/><path fill="currentColor" d="M18 6v.75h.75V6H18Zm-2.283-3.674l-.106-.742l.106.742ZM4.92 3.87l-.106-.743l.106.743Zm.15 2.88H18v-1.5H5.071v1.5ZM18.75 6V4.306h-1.5V6h1.5Zm-3.139-4.416L4.814 3.126l.212 1.485L15.823 3.07l-.212-1.485ZM4.814 3.126A1.821 1.821 0 0 0 3.25 4.93h1.5a.32.32 0 0 1 .276-.318l-.212-1.485Zm13.936 1.18a2.75 2.75 0 0 0-3.139-2.722l.212 1.485a1.25 1.25 0 0 1 1.427 1.237h1.5ZM5.071 5.25a.321.321 0 0 1-.321-.321h-1.5A1.82 1.82 0 0 0 5.071 6.75v-1.5Z"/></g>`),
		g.Group(children),
	)
}