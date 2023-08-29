package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandleB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20"/><path d="M24.792 23C27.668 23 30 20.761 30 18s-2.332-5-5.208-5H18v10h6.792Zm2.039 12C29.686 35 32 32.314 32 29s-2.314-6-5.169-6H18v12h8.831Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}