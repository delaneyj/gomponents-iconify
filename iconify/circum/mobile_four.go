package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="12.003" cy="18.937" r="1" fill="currentColor"/><path fill="currentColor" d="M16.725 2.065h-9.45a2.386 2.386 0 0 0-2.24 2.5v14.87a2.386 2.386 0 0 0 2.24 2.5h9.45a2.379 2.379 0 0 0 2.24-2.5V4.565a2.379 2.379 0 0 0-2.24-2.5Zm1.24 17.37a1.384 1.384 0 0 1-1.24 1.5h-9.45a1.391 1.391 0 0 1-1.24-1.5v-2.5h11.93Zm0-3.5H6.035V4.565a1.391 1.391 0 0 1 1.24-1.5h9.45a1.384 1.384 0 0 1 1.24 1.5Z"/>`),
		g.Group(children),
	)
}