package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreAddPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M9.5 9.5v-9h.25l.27.054a10.1 10.1 0 0 0 3.96 0L14.25.5h.25v9h9v.25l-.054.27a10.101 10.101 0 0 0 0 3.96l.054.27v.25h-9v9h-.25l-.27-.054a10.101 10.101 0 0 0-3.96 0l-.27.054H9.5v-9h-9v-.25l.054-.27a10.1 10.1 0 0 0 0-3.96L.5 9.75V9.5h9Z"/>`),
		g.Group(children),
	)
}