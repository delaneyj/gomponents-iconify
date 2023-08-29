package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpaceBetween(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 5h-4v14h4v-2h-2V7h2V5ZM5 5h4v14H5v-2h2V7H5V5Zm8 2v10h-2V7h2Z"/>`),
		g.Group(children),
	)
}