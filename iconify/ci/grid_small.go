package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 17h-4v-4h4v4Zm-6 0H7v-4h4v4Zm6-6h-4V7h4v4Zm-6 0H7V7h4v4Z"/>`),
		g.Group(children),
	)
}