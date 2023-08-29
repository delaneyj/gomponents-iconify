package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoadbarSound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 6h2v12h-2V6Zm-4 7h2v5H7v-5Zm8-4h2v9h-2V9Z"/>`),
		g.Group(children),
	)
}