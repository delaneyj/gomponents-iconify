package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoFramer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.086 1H19.5v8.5h-5.086l6.5 6.5H13v7.914l-8.5-8.5V7.5h5.086L3.086 1Zm9.328 6.5H17.5V3H7.914l4.5 4.5Zm-.828 2H6.5v5.086l4.5 4.5V14h5.086l-4.5-4.5Z"/>`),
		g.Group(children),
	)
}