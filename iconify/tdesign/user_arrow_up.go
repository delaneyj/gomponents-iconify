package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Zm12 5.094l4.914 4.858l-1.406 1.422L19 16.394v7.11h-2v-7.11l-2.508 2.48l-1.406-1.422L18 12.594ZM8 16a4 4 0 0 0-4 4h8.8v2H2v-2a6 6 0 0 1 6-6h4.75v2H8Z"/>`),
		g.Group(children),
	)
}