package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M22 8v12c0 2.21-4.03 4-9 4s-9-1.79-9-4V8"/><path d="M22 14c0 2.21-4.03 4-9 4s-9-1.79-9-4m18-6c0 2.21-4.03 4-9 4s-9-1.79-9-4s4.03-4 9-4s9 1.79 9 4Zm10-2h6a4 4 0 0 1 4 4v6M16 42h-6a4 4 0 0 1-4-4v-6"/><circle cx="35" cy="29" r="5"/><path d="M44 44H26a9 9 0 1 1 18 0Z"/></g>`),
		g.Group(children),
	)
}