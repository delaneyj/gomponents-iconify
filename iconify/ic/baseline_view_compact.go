package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineViewCompact(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18h2.5v-2.5H4V18zm0-4.75h2.5v-2.5H4v2.5zM4 8.5h2.5V6H4v2.5zM17.5 6v2.5H20V6h-2.5zM13 8.5h2.5V6H13v2.5zm4.5 9.5H20v-2.5h-2.5V18zm0-4.75H20v-2.5h-2.5v2.5zM8.5 18H11v-2.5H8.5V18zm4.5 0h2.5v-2.5H13V18zM8.5 8.5H11V6H8.5v2.5zm4.5 4.75h2.5v-2.5H13v2.5zm-4.5 0H11v-2.5H8.5v2.5z"/>`),
		g.Group(children),
	)
}