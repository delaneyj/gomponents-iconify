package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuUpOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 16v-1.5l-6-6l-6 6V16h12m-6-4.67L14.67 14H9.33L12 11.33Z"/>`),
		g.Group(children),
	)
}