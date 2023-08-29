package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.733 3.704A3.982 3.982 0 0 1 10 3h4a4 4 0 0 1 4 4v7m-.1 3.895A4 4 0 0 1 14 21h-4a4 4 0 0 1-4-4V7c0-.3.033-.593.096-.874M12 7v1M3 3l18 18"/>`),
		g.Group(children),
	)
}