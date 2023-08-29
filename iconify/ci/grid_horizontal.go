package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 17h-4v-4h4v4Zm-6 0h-4v-4h4v4Zm-6 0H4v-4h4v4Zm12-6h-4V7h4v4Zm-6 0h-4V7h4v4Zm-6 0H4V7h4v4Z"/>`),
		g.Group(children),
	)
}