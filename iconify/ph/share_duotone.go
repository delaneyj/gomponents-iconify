package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m224 104l-48 48V56Z" opacity=".2"/><path d="m229.66 98.34l-48-48A8 8 0 0 0 168 56v40h-3a103.94 103.94 0 0 0-100.75 78a8 8 0 0 0 15.5 4A88 88 0 0 1 165 112h3v40a8 8 0 0 0 13.66 5.66l48-48a8 8 0 0 0 0-11.32ZM184 132.69V75.31L212.69 104ZM200 216a8 8 0 0 1-8 8H40a16 16 0 0 1-16-16V88a8 8 0 0 1 16 0v120h152a8 8 0 0 1 8 8Z"/></g>`),
		g.Group(children),
	)
}