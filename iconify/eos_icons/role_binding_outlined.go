package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoleBindingOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 19a4 4 0 0 1-4 4h-2v-2h2a2 2 0 0 0 0-4h-2v-2h2a4 4 0 0 1 4 4ZM9 19a4 4 0 0 1 4-4h2v2h-2a2 2 0 0 0 0 4h2v2h-2a4 4 0 0 1-4-4Z"/><path fill="currentColor" d="M14 18h4v2h-4zM9 5a3 3 0 1 0 3 3a3.009 3.009 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1.003 1.003 0 0 1-1 1Zm-3.69 6A7.011 7.011 0 0 1 9 13.88a5.641 5.641 0 0 1 .778.064A5.965 5.965 0 0 1 13 13h.254A9.398 9.398 0 0 0 9 11.89c-2.03 0-6 1.07-6 3.58V17h4.349a5.986 5.986 0 0 1 1.188-2Z"/><path fill="currentColor" d="M16 2h-4.18a2.988 2.988 0 0 0-5.64 0H2a2.006 2.006 0 0 0-2 2v14a2.006 2.006 0 0 0 2 2h5.141a3.606 3.606 0 0 1 0-2H2V4h14v9h2V4a2.006 2.006 0 0 0-2-2ZM9 3.25a.755.755 0 0 1-.75-.75a.75.75 0 0 1 1.5 0a.755.755 0 0 1-.75.75Z"/>`),
		g.Group(children),
	)
}