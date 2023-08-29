package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RibbonSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 1a5 5 0 1 0 0 10A5 5 0 0 0 8 1ZM5 14.5v-3.303c.883.51 1.907.803 3 .803a5.972 5.972 0 0 0 3-.803V14.5a.5.5 0 0 1-.757.429L8 13.583L5.757 14.93A.5.5 0 0 1 5 14.5Z"/>`),
		g.Group(children),
	)
}